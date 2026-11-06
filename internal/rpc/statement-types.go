package rpc

type StatementRequest struct {
	WalletAddress string `json:"address"`
}

type StatementResponse struct {
	Balance      int64         `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}
