package rpc

func (r *BalanceRequest) NewRequest(walletAddress string) {
	r.WalletAddress = walletAddress
}

func (r *BalanceResponse) NewResponse(walletAddress string, balance int64) {
	r.WalletAddress = walletAddress
	r.Balance = balance
}
