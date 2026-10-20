package components

import (
	"bufio"
	"crypto/ecdsa"
	"errors"
	"fastcoin/src/handlers"
	"fastcoin/src/utils"
	"fmt"
	"os"
	"strings"
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
	err = c.create(keys)
	if err != nil {
		return err
	}
	err = c.listen()
	if err != nil {
		return err
	}
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

func (c *CoreNode) create(keys string) error {
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
	//blockchain handling
	c.Peers = make([]string, 0)
	fmt.Println("[INFO] Welcome to core node, type a command and it's parameters (type h for help).")
	return err
}

func (c *CoreNode) listen() error {
	rpc := RPCConnection{Type: "tcp", Address: c.CoreAddress}
	listener, err := rpc.Listen()
	if err != nil {
		fmt.Printf("[ERROR] Network Connection Error: core-%v\n", err)
		return err
	}
	go rpc.Receive(listener)
	return nil
}
