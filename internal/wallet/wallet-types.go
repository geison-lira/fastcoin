package wallet

import (
	"crypto/ecdsa"
)

type Node struct {
	PrivKey       *ecdsa.PrivateKey
	PubKey        *ecdsa.PublicKey
	WalletAddress string
	Nonce         int64
	Peers         []string
}
