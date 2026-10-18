package components

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Block struct {
	Id                 int64         `json:"id"`
	Timestamp          int64         `json:"timestamp"`
	Transactions       []Transaction `json:"transactions"`
	PrevHash           string        `json:"prevhash"`
	Hash               string        `json:"hash"`
	Nonce              int64         `json:"nonce"`
	MinerWalletAddress string        `json:"minerwalletaddress"`
}

func (b *Block) hashCalculate() {
	data := strconv.FormatInt(b.Id, 10) + strconv.FormatInt(b.Timestamp, 10)
	for _, tx := range b.Transactions {
		data += tx.Id
	}
	data += b.PrevHash + strconv.FormatInt(b.Nonce, 10) + b.MinerWalletAddress
	dataHash := sha256.Sum256([]byte(data))
	b.Hash = hex.EncodeToString(dataHash[:])
}
