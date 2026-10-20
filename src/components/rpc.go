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

type RPCResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type RPCConnection struct {
	Type    string
	Address string
	TimeOut time.Duration
}

func (r *RPCConnection) Send(req RPCRequest) (*RPCResponse, error) {
	conn, err := net.DialTimeout(r.Type, r.Address, r.TimeOut)
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

func (r *RPCConnection) Listen() (net.Listener, error) {
	listener, err := net.Listen(r.Type, r.Address)
	if err != nil {
		return nil, fmt.Errorf("unable to listen to the network at %s: %w", r.Address, err)
	}
	defer listener.Close()
	return listener, nil
}

func (r *RPCConnection) Receive(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		defer conn.Close()
		var req RPCRequest
		if err := json.NewDecoder(conn).Decode(&req); err != nil {
			return
		}
		var res RPCResponse
		// call the switch function
		if err := json.NewEncoder(conn).Encode(res); err != nil {
			return
		}
	}
}
