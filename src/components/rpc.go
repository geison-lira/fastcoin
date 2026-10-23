package components

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type RPCRequest struct {
	Method string `json:"method"`
	Params any    `json:"params"`
}

func (r *RPCRequest) NewRequest(method string, params any) {
	r.Method = method
	r.Params = params
}

type RPCResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func (r *RPCResponse) NewResponse(success bool, data any, err string) {
	r.Success = success
	r.Data = data
	r.Error = err
}

type RPCConnection struct {
	Protocol string
	Address  string
	TimeOut  time.Duration
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

func (r *RPCConnection) Handle(listenErrChan chan error, listenSigStop chan bool) error {
	listener, err := net.Listen(r.Protocol, r.Address)
	if err != nil {
		return fmt.Errorf("unable to listen to the network at %s: %w", r.Address, err)
	}
	defer listener.Close()
	go func() {
		for {
			select {
			case stop := <-listenSigStop:
				if stop {
					return
				}
			default:
			}
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			defer conn.Close()
			var req RPCRequest
			if err := json.NewDecoder(conn).Decode(&req); err != nil {
				listenErrChan <- err
				return
			}
			res := r.procedures(req)
			if err := json.NewEncoder(conn).Encode(res); err != nil {
				listenErrChan <- err
				return
			}
		}
	}()
	return err
}
