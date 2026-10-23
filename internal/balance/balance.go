package balance

func (r *BalanceRequest) NewRequest(walletAddress string) {
	r.WalletAddress = walletAddress
}
