package core

import "sync"

type CoreNode struct {
	Blockchain    *Blockchain
	Peers         []string
	CoreAddress   string
	WalletAddress string
	ListenSigStop chan bool
	Listening     bool
	MineSigStop   chan bool
	Mining        bool
	mu            sync.Mutex
}
