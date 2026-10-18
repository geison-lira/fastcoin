package components

import (
	"crypto/ecdsa"
	"errors"
	"fastcoin/src/handlers"
	"fastcoin/src/utils"
	"fmt"
	"sync"
)

const difficulty = 3

type CoreNode struct {
	Blockchain    *Blockchain
	Peers         []string
	CoreAddress   string
	WalletAddress string
	mu            sync.Mutex
}

func (c *CoreNode) Enter(keys string) error {
	var err error
	fmt.Println("[INFO] Launching core...")
	if !utils.FileExists(keys) {
		err = errors.New("missing miner wallet keys")
		fmt.Printf("[FATAL] Keys Inexistence Error: core-%v\n", err)
		return err
	}
	err = c.load(keys)
	if err != nil {
		return err
	}
}

func (c *CoreNode) load(keys string) error {
	var err error
	var pubKey *ecdsa.PublicKey
	fmt.Println("[INFO] Keys found, loading core...")
	_, pubKey, err = handlers.KeysLoad(keys)
	if err != nil {
		fmt.Printf("[FATAL] Key Load Error: core-%v\n", err)
		return err
	}
	c.WalletAddress, err = handlers.AddressGenerate(pubKey)
	if err != nil {
		fmt.Printf("[FATAL] Address Generation Error: core-%v\n", err)
		return err
	}
	c.Blockchain.blockchainGenerate(difficulty)
	c.Peers = make([]string, 0)
	fmt.Println("[INFO] Welcome to core node, type a command and it's parameters (type h for help).")
	return err
}
