package rpc

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func (t *Transaction) NewTransaction(sender string, recipient string, ammount int64, nonce int64, privKey *ecdsa.PrivateKey, pubKey string) error {
	var err error
	t.Sender = sender
	t.Recipient = recipient
	t.Ammount = ammount
	t.Nonce = nonce
	t.SenderPubKey = pubKey
	t.Id = t.Hash()
	t.Signature, err = t.sign(privKey)
	return err
}

func (t *Transaction) NewReward(sender string, recipient string, ammount int64, nonce int64) {
	t.Sender = sender
	t.Recipient = recipient
	t.Ammount = ammount
	t.Nonce = nonce
	t.Id = t.Hash()
}

func (t *Transaction) Hash() string {
	txFields := fmt.Sprintf("%s:%s:%d:%d", t.Sender, t.Recipient, t.Ammount, t.Nonce)
	hashTxFields := sha256.Sum256([]byte(txFields))
	return hex.EncodeToString(hashTxFields[:])
}

func (t *Transaction) sign(privKey *ecdsa.PrivateKey) (string, error) {
	signBytes, err := ecdsa.SignASN1(rand.Reader, privKey, []byte(t.Id))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signBytes), nil
}
