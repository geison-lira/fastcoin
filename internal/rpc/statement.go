package rpc

func (r *StatementRequest) NewRequest(walletAddress string) {
	r.WalletAddress = walletAddress
}

func (r *StatementResponse) NewResponse(balance int64, transactions []Transaction) {
	r.Balance = balance
	r.Transactions = transactions
}
