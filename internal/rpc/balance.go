package rpc

func (r *BalanceRequest) NewRequest(walletAddress string) {
	r.WalletAddress = walletAddress
}
