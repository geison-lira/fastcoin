package rpc

func (r *StatementRequest) NewRequest(walletAddress string) {
	r.WalletAddress = walletAddress
}

func (r *StatementResponse) NewResponse(walletAddress string, balance int64, transactions []Transaction) {
	r.WalletAddress = walletAddress
	r.Balance = balance
	r.Transactions = transactions
}
