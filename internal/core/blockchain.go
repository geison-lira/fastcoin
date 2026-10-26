package core

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fastcoin/internal/rpc"
	"strconv"
	"strings"
	"time"
)

const reward = 50

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

func (b *Block) updateHash() {
	b.Nonce++
	b.Hash = b.hash()
}

func (bc *Blockchain) newBlockchain(difficulty int) {
	bc.Blocks = make([]Block, 0)
	bc.Mempool = make([]rpc.Transaction, 0)
	bc.Balances = make(map[string]int64)
	bc.Nonces = make(map[string]int64)
	bc.Difficulty = difficulty
	var zeroBlock Block
	zeroBlock.newBlock(0, nil, "0000000000000000000000000000000000000000000000000000000000000000", "SYSTEM")
	bc.Blocks = append(bc.Blocks, zeroBlock)
}

func (bc *Blockchain) mine(miner string, mineSigStop chan bool) {
	go func() {
		for {
			select {
			case <-mineSigStop:
				return
			default:
			}
			bc.mu.Lock()
			var rewardTx rpc.Transaction
			prevBlock := bc.Blocks[len(bc.Blocks)-1]
			nonce := time.Now().Unix()
			rewardTx.NewReward("SYSTEM", miner, reward, nonce)
			txs := append([]rpc.Transaction{rewardTx}, bc.Mempool...)
			bc.Mempool = make([]rpc.Transaction, 0)
			bc.mu.Unlock()
			var currBlock Block
			currBlock.newBlock(prevBlock.Id+1, txs, prevBlock.Hash, miner)
			target := strings.Repeat("0", bc.Difficulty)
			for {
				if strings.HasPrefix(currBlock.Hash, target) {
					break
				}
				currBlock.updateHash()
			}
			bc.mu.Lock()
			bc.Blocks = append(bc.Blocks, currBlock)
			bc.updateBalances(currBlock)
			bc.mu.Unlock()
		}
	}()
}

func (bc *Blockchain) updateBalances(currBlock Block) {
	for _, tx := range currBlock.Transactions {
		if tx.Sender != "SYSTEM" {
			bc.Balances[tx.Sender] -= tx.Ammount
		}
		bc.Balances[tx.Recipient] += tx.Ammount
	}
}

func (bc *Blockchain) FetchStatement(walletAddress string) (int64, []rpc.Transaction, error) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	var txs []rpc.Transaction
	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			if tx.Sender == walletAddress || tx.Recipient == walletAddress {
				txs = append(txs, tx)
			}
		}
	}
	bl, exists := bc.Balances[walletAddress]
	if !exists {
		err := errors.New("wallet does not exist")
		return 0, nil, err
	}
	return bl, txs, nil
}

func (bc *Blockchain) FetchBalance(walletAddress string) (int64, error) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	bl, exists := bc.Balances[walletAddress]
	if !exists {
		err := errors.New("wallet does not exist")
		return 0, err
	}
	return bl, nil
}
