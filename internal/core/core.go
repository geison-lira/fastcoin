package core

import (
	"bufio"
	"crypto/ecdsa"
	"fastcoin/internal/handlers"
	"fmt"
	"os"
	"strings"
)

func (c *CoreNode) NewCoreNode(coreAddress string, keys string) error {
	var err error
	var pubKey *ecdsa.PublicKey
	fmt.Println("[INFO] Keys found, creating core...")
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
	//blockchain creation and handling here
	c.CoreAddress = coreAddress
	c.Peers = make([]string, 0)
	c.ListenErrChan = make(chan error)
	c.ListenSigStop = make(chan bool)
	fmt.Println("[INFO] Welcome to core node, type a command and it's parameters (type h for help).")
	return err
}

func (c *CoreNode) Enter(keys string) error {
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("fastcoin-core@%s...> ", c.WalletAddress[:8])
		if !scanner.Scan() {
			err = scanner.Err()
			if err != nil {
				fmt.Printf("[FATAL] Input Read Error: core-%v\n", err)
				return err
			}
			fmt.Println("[INFO] Closing core...")
			os.Exit(0)
		}
		usrInput := strings.Fields(scanner.Text())
		if len(usrInput) == 0 {
			continue
		}
		exit := c.control(usrInput)
		if exit {
			return err
		}
	}
}
