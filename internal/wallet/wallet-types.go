package wallet

import (
	"crypto/ecdsa"
)

type WalletNode struct {
	PrivKey       *ecdsa.PrivateKey
	PubKey        *ecdsa.PublicKey
	WalletAddress string
	Nonce         int64
	Peers         []string
}
