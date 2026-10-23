package wallet

import "crypto/ecdsa"

type WalletNode struct {
	PrivKey       *ecdsa.PrivateKey
	PubKey        *ecdsa.PublicKey
	WalletAddress string
	CoreAddress   string
	Nonce         int64
}
