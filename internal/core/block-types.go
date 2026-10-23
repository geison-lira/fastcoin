package core

import "fastcoin/internal/rpc"

type Block struct {
	Id                 int64             `json:"id"`
	Timestamp          int64             `json:"timestamp"`
	Transactions       []rpc.Transaction `json:"transactions"`
	PrevHash           string            `json:"prevhash"`
	Hash               string            `json:"hash"`
	Nonce              int64             `json:"nonce"`
	MinerWalletAddress string            `json:"minerwalletaddress"`
}
