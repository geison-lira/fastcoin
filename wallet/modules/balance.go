package modules

type BalanceRequest struct {
	Address string `json:"address"`
}

type BalanceResponse struct {
	Address string `json:"address"`
	Balance int64  `json:"balance"`
}
