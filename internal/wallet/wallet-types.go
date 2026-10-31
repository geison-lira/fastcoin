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
	Peers         []string
	Connection    rpc.RPCConnection
}
