type RPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:params`
}

type RPCResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}