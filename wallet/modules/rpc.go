package types

type RPCRequest struct {
	Method string `json:"method"`
	Params any    `json:"params"`
}

type RPCResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
