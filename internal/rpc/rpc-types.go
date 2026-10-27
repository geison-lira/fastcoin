package rpc

import (
	"encoding/json"
	"time"
)

type RPCRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

type RPCResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type RPCConnection struct {
	Protocol string
	Address  string
	TimeOut  time.Duration
}
