package core

import (
	"fastcoin/internal/rpc"
	"sync"
)

type Block struct {
	Id           int64             `json:"id"`
	Timestamp    int64             `json:"timestamp"`
	Transactions []rpc.Transaction `json:"transactions"`
	PrevHash     string            `json:"prevhash"`
	Hash         string            `json:"hash"`
	Nonce        int64             `json:"nonce"`
	Miner        string            `json:"miner"`
}

type Blockchain struct {
	Blocks     []Block
	Mempool    []rpc.Transaction
	Balances   map[string]int64
	Difficulty int
	mu         sync.RWMutex
}
