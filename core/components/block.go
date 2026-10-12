package components

type Block struct {
	Id            int64         `json:"id"`
	Timestamp     int64         `json:"timestamp"`
	Transactions  []Transaction `json:"transactions"`
	PrevHash      string        `json:"prevhash"`
	Hash          string        `json:"hash"`
	Nonce         int64         `json:"nonce"`
	WalletAddress string        `json:"walletaddress"`
}
