package components

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Transaction struct {
	Id        string `json:"id"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Ammount   int64  `json:"ammount"`
	Nonce     int64  `json:"nonce"`
	Signature string `json:"signature"`
}

func (t *Transaction) NewTransaction(sender string, recipient string, ammount int64, nonce int64) {
	t.Sender = sender
	t.Recipient = recipient
	t.Ammount = ammount
	t.Nonce = nonce
	t.Id = t.hash()
}
func (t *Transaction) Sign(signBytes []byte) {
	t.Signature = hex.EncodeToString(signBytes)
}

func (t *Transaction) hash() string {
	txFields := fmt.Sprintf("%s:%s:%d:%d", t.Sender, t.Recipient, t.Ammount, t.Nonce)
	hashTxFields := sha256.Sum256([]byte(txFields))
	return hex.EncodeToString(hashTxFields[:])
}
