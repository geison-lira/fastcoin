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
		switch usrInput[0] {
		case "mn":
			//mine
		case "h":
			if len(usrInput) == 1 {
				c.help()
			} else {
				fmt.Println("[ERROR] Command Typing Error: core-wrong command syntax (type h for help).")
			}
		case "q":
			if len(usrInput) == 1 {
				fmt.Println("[INFO] Exiting core...")
				return err
			} else {
				fmt.Println("[ERROR] Command Typing Error: core-wrong command syntax (type h for help).")
			}
		default:
			fmt.Println("[ERROR] Command Typing Error: core-unknown command (type h for help).")
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
	c.Blockchain.blockchainGenerate(difficulty)
	c.Peers = make([]string, 0)
	fmt.Println("[INFO] Welcome to core node, type a command and it's parameters (type h for help).")
	return err
}

func (c *CoreNode) help() {
	fmt.Println("      Command      |            Structure            ")
	fmt.Println("Command List Help  |  h                              ")
	fmt.Println("Quit Program       |  q                              ")
}
