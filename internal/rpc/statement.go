package rpc

func (r *StatementRequest) NewRequest(walletAddress string) {
	r.WalletAddress = walletAddress
}
