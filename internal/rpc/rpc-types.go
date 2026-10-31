package rpc

import (
	"encoding/json"
	"time"
)

type RPCRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params,omitempty"`
}

type RPCResponse struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   string          `json:"error,omitempty"`
}

type RPCConnection struct {
	Protocol string
	Address  string
	TimeOut  time.Duration
}
