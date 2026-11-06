package wallet

import (
	"bufio"
	"errors"
	"fastcoin/internal/handlers"
	"fastcoin/internal/utils"
	"fmt"
	"os"
	"strings"
	"time"
)

func (w *WalletNode) NewWalletNode(peers string, keys string) error {
	var err error
	if !utils.FileExists(peers) {
		err = errors.New("missing peer nodes file")
		fmt.Printf("[INIT-FATAL] Peers Inexistence Error: init-%v\n", err)
		return err
	}
	w.Peers, err = handlers.PeersLoad(peers)
	if err != nil {
		fmt.Printf("[INIT-FATAL] Peer Load Error: init-%v\n", err)
		return err
	}
	if len(w.Peers) == 0 {
		err = errors.New("no peers detected")
		fmt.Printf("[INIT-FATAL] Peer Selection Error: wallet-%v\n", err)
		return err
	}
	if !utils.FileExists(keys) {
		fmt.Printf("[INIT-INFO] No keys found, generate new keys? [y/n] ")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			if err = scanner.Err(); err != nil {
				fmt.Printf("[INIT-FATAL] Input Read Error: init-%v\n", err)
				return err
			}
			fmt.Println("[INIT-INFO] Closing app...")
			os.Exit(0)
		}
		usrInput := strings.Fields(scanner.Text())
		if !strings.EqualFold(usrInput[0], "y") {
			fmt.Println("[INIT-INFO] Closing app...")
			os.Exit(0)
		}
		w.PrivKey, w.PubKey, err = handlers.KeysGenerate()
		if err != nil {
			fmt.Printf("[INIT-FATAL] Key Generation Error: init-%v\n", err)
			return err
		}
		err = handlers.KeysStore(keys, w.PrivKey, w.PubKey)
		if err != nil {
			fmt.Printf("[INIT-FATAL] Key Storage Error: init-%v\n", err)
			return err
		}
	} else {
		fmt.Println("[INIT-INFO] Keys found, loading wallet.")
		w.PrivKey, w.PubKey, err = handlers.KeysLoad(keys)
		if err != nil {
			fmt.Printf("[INIT-FATAL] Key Load Error: init-%v\n", err)
			return err
		}
	}
	w.WalletAddress, err = handlers.AddressGenerate(w.PubKey)
	if err != nil {
		fmt.Printf("[INIT-FATAL] Address Generation Error: init-%v\n", err)
		return err
	}
	w.Nonce = time.Now().UnixNano()
	fmt.Println("[INIT-SUCCESS] Welcome to your FastCoin wallet, your keys are stored in keys.pem, type h to see commands.")
	return nil
}

func (w *WalletNode) Enter() error {
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("fastcoin-wallet@%s...> ", w.WalletAddress[:8])
		if !scanner.Scan() {
			if err = scanner.Err(); err != nil {
				fmt.Printf("[WALLET-FATAL] Input Read Error: wallet-%v\n", err)
				return err
			}
			fmt.Println("[WALLET-INFO] Closing wallet...")
			return nil
		}
		usrInput := strings.Fields(scanner.Text())
		if len(usrInput) == 0 {
			continue
		}
		exit := w.control(usrInput)
		if exit {
			return nil
		}
	}
}
