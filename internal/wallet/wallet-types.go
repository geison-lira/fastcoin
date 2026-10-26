package wallet

import (
	"crypto/ecdsa"
	"fastcoin/internal/rpc"
)

type WalletNode struct {
	PrivKey       *ecdsa.PrivateKey
	PubKey        *ecdsa.PublicKey
	WalletAddress string
	CoreAddress   string
	Nonce         int64
	Connection    rpc.RPCConnection
}
