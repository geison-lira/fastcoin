package core

import "sync"

type Blockchain struct {
	Blocks     []Block
	Mempool    []Transaction
	Balances   map[string]int64
	Difficulty int
	mu         sync.RWMutex
}
