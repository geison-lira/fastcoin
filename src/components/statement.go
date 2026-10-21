package components

type StatementRequest struct {
	WalletAddress string `json:"address"`
}

func (r *StatementRequest) NewRequest(walletAddress string) {
	r.WalletAddress = walletAddress
}

type StatementResponse struct {
	WalletAddress string        `json:"address"`
	Balance       int64         `json:"balance"`
	Transactions  []Transaction `json:"transactions"`
}
