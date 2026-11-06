package rpc

type BalanceRequest struct {
	WalletAddress string `json:"address"`
}

type BalanceResponse struct {
	Balance int64 `json:"balance"`
}
