package rpc

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func (r *RPCRequest) NewRequest(method string, params any) {
	r.Method = method
	r.Params = params
}

func (r *RPCResponse) NewResponse(success bool, data any, err string) {
	r.Success = success
	r.Data = data
	r.Error = err
}

func (r *RPCConnection) NewConnection(protocol string, address string, timeOut time.Duration) {
	r.Protocol = protocol
	r.Address = address
	r.TimeOut = timeOut
}

func (r *RPCConnection) Send(req RPCRequest) (*RPCResponse, error) {
	conn, err := net.DialTimeout(r.Protocol, r.Address, r.TimeOut)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the network at %s: %w", r.Address, err)
	}
	defer conn.Close()
	if err := json.NewEncoder(conn).Encode(req); err != nil {
		return nil, err
	}
	var res RPCResponse
	if err := json.NewDecoder(conn).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *RPCConnection) Handle(listenSigStop chan bool) error {
	listener, err := net.Listen(r.Protocol, r.Address)
	if err != nil {
		return fmt.Errorf("unable to listen to the network at %s: %w", r.Address, err)
	}
	go func() {
		defer listener.Close()
		<-listenSigStop
	}()
	go func() {
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go func(conn net.Conn) {
				defer conn.Close()
				var req RPCRequest
				if err := json.NewDecoder(conn).Decode(&req); err != nil {
					return
				}
				res := r.procedures(req)
				if err := json.NewEncoder(conn).Encode(res); err != nil {
					return
				}
			}(conn)
		}
	}()
	return nil
}
