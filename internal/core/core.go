package core

import (
	"bufio"
	"crypto/ecdsa"
	"fastcoin/internal/handlers"
	"fmt"
	"os"
	"strings"
)

const difficulty = 3

func (c *CoreNode) NewCoreNode(coreAddress string, keys string) error {
	var err error
	var pubKey *ecdsa.PublicKey
	fmt.Println("[INIT-INFO] Keys found, loading core.")
	_, pubKey, err = handlers.KeysLoad(keys)
	if err != nil {
		fmt.Printf("[INIT-FATAL] Key Load Error: init-%v\n", err)
		return err
	}
	c.WalletAddress, err = handlers.AddressGenerate(pubKey)
	if err != nil {
		fmt.Printf("[INIT-FATAL] Address Generation Error: init-%v\n", err)
		return err
	}
	c.Blockchain.newBlockchain(difficulty)
	c.CoreAddress = coreAddress
	c.Peers = make([]string, 0)
	c.ListenSigStop = make(chan bool)
	fmt.Println("[INIT-SUCCESS] Welcome to your new FastCoin core, type h to see commands.")
	return nil
}

func (c *CoreNode) Enter(keys string) error {
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("fastcoin-core@%s...> ", c.WalletAddress[:8])
		if !scanner.Scan() {
			err = scanner.Err()
			if err != nil {
				fmt.Printf("[CORE-FATAL] Input Read Error: core-%v\n", err)
				return err
			}
			fmt.Println("[CORE-INFO] Closing core...")
			return nil
		}
		usrInput := strings.Fields(scanner.Text())
		if len(usrInput) == 0 {
			continue
		}
		exit := c.control(usrInput)
		if exit {
			return nil
		}
	}
}
