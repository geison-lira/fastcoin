package components

type BalanceRequest struct {
	WalletAddress string `json:"address"`
}

type BalanceResponse struct {
	WalletAddress string `json:"address"`
	Balance       int64  `json:"balance"`
}
