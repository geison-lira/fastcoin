package components

import "sync"

type Blockchain struct {
	mu         sync.RWMutex
	Blocks     []Block
	Mempool    []Transaction
	Balances   map[string]int64
	Difficulty int
}
