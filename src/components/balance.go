package components

type BalanceRequest struct {
	WalletAddress string `json:"address"`
}

func (r *BalanceRequest) NewRequest(walletAddress string) {
	r.WalletAddress = walletAddress
}

type BalanceResponse struct {
	WalletAddress string `json:"address"`
	Balance       int64  `json:"balance"`
}
