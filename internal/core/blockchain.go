package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fastcoin/internal/rpc"
	"strconv"
	"strings"
	"time"
)

func (b *Block) newBlock(id int64, transactions []rpc.Transaction, prevHash string, miner string) {
	b.Id = id
	b.Timestamp = time.Now().Unix()
	b.Transactions = transactions
	b.PrevHash = prevHash
	b.Miner = miner
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
	zeroBlock.newBlock(0, nil, "0000000000000000000000000000000000000000000000000000000000000000", "SYSTEM")
	bc.Blocks = append(bc.Blocks, zeroBlock)
}
