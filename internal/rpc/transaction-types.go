package rpc

type Transaction struct {
	Id           string `json:"id"`
	Sender       string `json:"sender"`
	Recipient    string `json:"recipient"`
	Ammount      int64  `json:"ammount"`
	Nonce        int64  `json:"nonce"`
	SenderPubKey string `json:"pubkey"`
	Signature    string `json:"signature"`
}
