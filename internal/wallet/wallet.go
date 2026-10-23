package wallet

import (
	"bufio"
	"fastcoin/internal/handlers"
	"fmt"
	"os"
	"strings"
	"time"
)

func (w *WalletNode) NewWalletNode(coreAddress string, keys string) error {
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
	w.CoreAddress = coreAddress
	w.Nonce = time.Now().UnixNano()
	fmt.Println("[INFO] Welcome to your new FastCoin wallet, your wallet was successfully created and your keys are stored in keys.pem.")
	return err
}

func (w *WalletNode) LoadWalletNode(coreAddress string, keys string) error {
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
	w.CoreAddress = coreAddress
	w.Nonce = time.Now().UnixNano()
	fmt.Println("[INFO] Welcome back to your wallet, type a command and it's parameters (type h for help).")
	return err
}

func (w *WalletNode) Enter(keys string) error {
	var err error
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
			return err
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
