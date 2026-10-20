package components

import (
	"bufio"
	"crypto/ecdsa"
	"fastcoin/src/handlers"
	"fastcoin/src/utils"
	"fmt"
	"os"
	"strings"
)

type WalletNode struct {
	PrivKey       *ecdsa.PrivateKey
	PubKey        *ecdsa.PublicKey
	WalletAddress string
	CoreAddress   string
	Nonce         int64
}

func (w *WalletNode) Enter(keys string) error {
	var err error
	fmt.Println("[INFO] Launching wallet...")
	if !utils.FileExists(keys) {
		err = w.create(keys)
		if err != nil {
			return err
		}
	} else {
		err = w.load(keys)
		if err != nil {
			return err
		}
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("fastcoin-wallet@%s...> ", w.WalletAddress[:8])
		if !scanner.Scan() {
			err = scanner.Err()
			if err != nil {
				fmt.Printf("[FATAL] Input Read Error: wallet-%v\n", err)
				return err
			}
			fmt.Println("[INFO] Closing wallet...")
			os.Exit(0)
		}
		usrInput := strings.Fields(scanner.Text())
		if len(usrInput) == 0 {
			continue
		}
		exit := w.control(usrInput)
		if exit {
			return err
		}
	}
}

func (w *WalletNode) create(keys string) error {
	var err error
	fmt.Println("[INFO] No keys found, generate new ones? [y/n]")
	fmt.Printf("fastcoin-wallet@root> ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err = scanner.Err(); err != nil {
			fmt.Printf("[FATAL] Input Read Error: wallet-%v\n", err)
			return err
		}
		fmt.Println("[INFO] Closing wallet...")
		os.Exit(0)
	}
	usrInput := strings.Fields(scanner.Text())
	if !strings.EqualFold(usrInput[0], "y") {
		fmt.Println("[INFO] Closing wallet...")
		os.Exit(0)
	}
	w.PrivKey, w.PubKey, err = handlers.KeysGenerate()
	if err != nil {
		fmt.Printf("[FATAL] Key Generation Error: wallet-%v\n", err)
		return err
	}
	err = handlers.KeysStore(keys, w.PrivKey, w.PubKey)
	if err != nil {
		fmt.Printf("[FATAL] Key Storage Error: wallet-%v\n", err)
		return err
	}
	w.WalletAddress, err = handlers.AddressGenerate(w.PubKey)
	if err != nil {
		fmt.Printf("[FATAL] Address Generation Error: wallet-%v\n", err)
		return err
	}
	fmt.Println("[INFO] Welcome to your new FastCoin wallet, your wallet was successfully created and your keys are stored in keys.pem.")
	return err
}

func (w *WalletNode) load(keys string) error {
	var err error
	fmt.Println("[INFO] Keys found, loading wallet...")
	w.PrivKey, w.PubKey, err = handlers.KeysLoad(keys)
	if err != nil {
		fmt.Printf("[FATAL] Key Load Error: wallet-%v\n", err)
		return err
	}
	w.WalletAddress, err = handlers.AddressGenerate(w.PubKey)
	if err != nil {
		fmt.Printf("[FATAL] Address Generation Error: wallet-%v\n", err)
		return err
	}
	fmt.Println("[INFO] Welcome back to your wallet, type a command and it's parameters (type h for help).")
	return err
}
