package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fastcoin/internal/rpc"
	"strconv"
	"strings"
	"time"
)

func (b *Block) newBlock() {
	b.Id = 0
	b.Timestamp = time.Now().Unix()
	b.Transactions = make([]rpc.Transaction, 0)
	b.PrevHash = "0000000000000000000000000000000000000000000000000000000000000000"
	b.Nonce = 0
	b.Miner = "SYSTEM"
	b.Hash = b.hash()
}

func (b *Block) hash() string {
	var builder strings.Builder
	data := strconv.FormatInt(b.Id, 10) + strconv.FormatInt(b.Timestamp, 10)
	for _, tx := range b.Transactions {
		builder.WriteString(tx.Id)
	}
	data += builder.String()
	data += b.PrevHash + strconv.FormatInt(b.Nonce, 10) + b.Miner
	dataHash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(dataHash[:])
}

func (bc *Blockchain) newBlockchain(difficulty int) {
	bc.Blocks = make([]Block, 0)
	bc.Mempool = make([]rpc.Transaction, 0)
	bc.Balances = make(map[string]int64)
	bc.Difficulty = difficulty
	var zeroBlock Block
	zeroBlock.newBlock()
	bc.Blocks = append(bc.Blocks, zeroBlock)
}
