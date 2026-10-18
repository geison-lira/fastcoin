package components

import (
	"bufio"
	"crypto/ecdsa"
	"fastcoin/src/handlers"
	"fastcoin/src/utils"
	"fmt"
	"os"
	"strings"
	"sync"
)

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
		err = c.create(keys)
		if err != nil {
			return err
		}
	} else {
		err = c.load(keys)
		if err != nil {
			return err
		}
	}
}

func (c *CoreNode) create(keys string) error {
	var err error
	var privKey *ecdsa.PrivateKey
	var pubKey *ecdsa.PublicKey
	fmt.Println("[INFO] No keys found, generate new ones? [y/n]")
	fmt.Printf("fastcoin-core@root> ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err = scanner.Err(); err != nil {
			fmt.Printf("[FATAL] Input Read Error: core-%v\n", err)
			return err
		}
		fmt.Println("[INFO] Closing core...")
		os.Exit(0)
	}

	usrInput := strings.Fields(scanner.Text())
	if !strings.EqualFold(usrInput[0], "y") {
		fmt.Println("[INFO] Closing core...")
		os.Exit(0)
	}
	privKey, pubKey, err = handlers.KeysGenerate()
	if err != nil {
		fmt.Printf("[FATAL] Key Generation Error: core-%v\n", err)
		return err
	}
	err = handlers.KeysStore(keys, privKey, pubKey)
	if err != nil {
		fmt.Printf("[FATAL] Key Storage Error: core-%v\n", err)
		return err
	}
	c.WalletAddress, err = handlers.AddressGenerate(pubKey)
	if err != nil {
		fmt.Printf("[FATAL] Address Generation Error: core-%v\n", err)
		return err
	}
	c.Peers = make([]string, 0)
	fmt.Println("[INFO] Welcome to your new FastCoin core node, the associated wallet was successfully created and the keys are stored in keys.pem.")
	return err
}
