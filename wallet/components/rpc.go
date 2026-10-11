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

func (r *RPCConnection) send(req RPCRequest) (*RPCResponse, error) {
	conn, err := net.DialTimeout(r.Type, r.Address, r.TimeOut)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the network at %s: %w", r.Address, err)
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
