package core

import (
	"time"
)

func (bc *Blockchain) blockchainGenerate(difficulty int) {
	bc.Blocks = make([]Block, 0)
	bc.Mempool = make([]Transaction, 0)
	bc.Balances = make(map[string]int64)
	bc.Difficulty = difficulty
	genesisBlock := Block{
		Id:                 0,
		Timestamp:          time.Now().Unix(),
		Transactions:       make([]Transaction, 0),
		PrevHash:           "0000000000000000000000000000000000000000000000000000000000000000",
		Nonce:              0,
		MinerWalletAddress: "GENESIS",
	}
	genesisBlock.hashCalculate()
	bc.Blocks = append(bc.Blocks, genesisBlock)
}
