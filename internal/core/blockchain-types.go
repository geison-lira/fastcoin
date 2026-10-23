package core

import (
	"fastcoin/internal/rpc"
	"sync"
)

type Blockchain struct {
	Blocks     []Block
	Mempool    []rpc.Transaction
	Balances   map[string]int64
	Difficulty int
	mu         sync.RWMutex
}
