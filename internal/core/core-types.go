package core

import "sync"

type CoreNode struct {
	Blockchain    *Blockchain
	Peers         []string
	CoreAddress   string
	WalletAddress string
	ListenErrChan chan error
	ListenSigStop chan bool
	Listening     bool
	mu            sync.Mutex
}
