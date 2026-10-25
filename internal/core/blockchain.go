package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fastcoin/internal/rpc"
	"strconv"
	"strings"
	"time"
)

func (b *Block) hashCalculate() {
	var builder strings.Builder
	data := strconv.FormatInt(b.Id, 10) + strconv.FormatInt(b.Timestamp, 10)
	for _, tx := range b.Transactions {
		builder.WriteString(tx.Id)
	}
	data += builder.String()
	data += b.PrevHash + strconv.FormatInt(b.Nonce, 10) + b.MinerWalletAddress
	dataHash := sha256.Sum256([]byte(data))
	b.Hash = hex.EncodeToString(dataHash[:])
}

func (bc *Blockchain) blockchainGenerate(difficulty int) {
	bc.Blocks = make([]Block, 0)
	bc.Mempool = make([]rpc.Transaction, 0)
	bc.Balances = make(map[string]int64)
	bc.Difficulty = difficulty
	genesisBlock := Block{
		Id:                 0,
		Timestamp:          time.Now().Unix(),
		Transactions:       make([]rpc.Transaction, 0),
		PrevHash:           "0000000000000000000000000000000000000000000000000000000000000000",
		Nonce:              0,
		MinerWalletAddress: "GENESIS",
	}
	genesisBlock.hashCalculate()
	bc.Blocks = append(bc.Blocks, genesisBlock)
}
