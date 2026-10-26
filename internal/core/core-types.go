package core

import (
	"fastcoin/internal/rpc"
	"sync"
)

type CoreNode struct {
	Blockchain    *Blockchain
	Peers         []string
	CoreAddress   string
	WalletAddress string
	ListenSigStop chan bool
	Listening     bool
	MineSigStop   chan bool
	Mining        bool
	Connection    rpc.RPCConnection
	mu            sync.Mutex
}
