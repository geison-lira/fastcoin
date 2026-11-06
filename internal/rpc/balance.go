package rpc

func (r *BalanceRequest) NewRequest(walletAddress string) {
	r.WalletAddress = walletAddress
}

func (r *BalanceResponse) NewResponse(balance int64) {
	r.Balance = balance
}
