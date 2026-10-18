package components

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fastcoin/src/handlers"
	"fastcoin/src/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
		switch usrInput[0] {
		case "tx":
			if len(usrInput) == 3 {
				w.transaction(usrInput[1:])
			} else {
				fmt.Println("[ERROR] Command Typing Error: wallet-wrong command syntax (type h for help).")
			}
		case "st":
			if len(usrInput) == 1 {
				w.statement()
			} else {
				fmt.Println("[ERROR] Command Typing Error: wallet-wrong command syntax (type h for help).")
			}
		case "bl":
			if len(usrInput) == 1 {
				w.balance()
			} else {
				fmt.Println("[ERROR] Command Typing Error: wallet-wrong command syntax (type h for help).")
			}
		case "h":
			if len(usrInput) == 1 {
				w.help()
			} else {
				fmt.Println("[ERROR] Command Typing Error: wallet-wrong command syntax (type h for help).")
			}
		case "q":
			if len(usrInput) == 1 {
				fmt.Println("[INFO] Exiting wallet...")
				return err
			} else {
				fmt.Println("[ERROR] Command Typing Error: wallet-wrong command syntax (type h for help).")
			}
		default:
			fmt.Println("[ERROR] Command Typing Error: wallet-unknown command (type h for help).")
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

func (w *WalletNode) transaction(usrInput []string) {
	ammount, err := strconv.ParseInt(usrInput[0], 10, 64)
	if err != nil {
		fmt.Printf("[ERROR] Transaction Parsing Error: wallet-%v\n", err)
		return
	}
	if ammount <= 0 {
		fmt.Println("[ERROR] Transaction Ammount Error: wallet-invalid ammount")
		return
	}
	recipient := usrInput[1]
	w.Nonce++
	tx := Transaction{Sender: w.WalletAddress, Recipient: recipient, Ammount: ammount, Nonce: w.Nonce}
	tx.Id = tx.Hash()
	signBytes, err := ecdsa.SignASN1(rand.Reader, w.PrivKey, []byte(tx.Id))
	if err != nil {
		fmt.Printf("[ERROR] Transaction Signing Error: wallet-%v\n", err)
		return
	}
	tx.Signature = hex.EncodeToString(signBytes)
	req := RPCRequest{Method: "set_tx", Params: tx}
	rpc := RPCConnection{Type: "tcp", Destination: w.CoreAddress, TimeOut: 3 * time.Second}
	res, err := rpc.Send(req)
	if err != nil {
		fmt.Printf("[ERROR] Network Connection Error: wallet-%v\n", err)
		return
	}
	if !res.Success {
		fmt.Printf("[ERROR] Transaction Broadcast Error: core-%s\n", res.Error)
		return
	}
	fmt.Println("[SUCCESS] Transaction Broadcast Successful")
	fmt.Printf("TxId: %s\n", tx.Id)
}

func (w *WalletNode) statement() {
	stReq := StatementRequest{WalletAddress: w.WalletAddress}
	req := RPCRequest{Method: "get_st", Params: stReq}
	rpc := RPCConnection{Type: "tcp", Destination: w.CoreAddress, TimeOut: 3 * time.Second}
	res, err := rpc.Send(req)
	if err != nil {
		fmt.Printf("[ERROR] Network Connection Error: wallet-%v\n", err)
		return
	}
	if !res.Success {
		fmt.Printf("[ERROR] Statement Fetching Error: core-%s\n", res.Error)
		return
	}
	bytes, _ := json.Marshal(res.Data)
	var stRes StatementResponse
	json.Unmarshal(bytes, &stRes)
	fmt.Println("[SUCCESS] Statement Fetching Successful")
	fmt.Printf("Wallet:\n%s\nBalance:\n%d FST\n", w.WalletAddress, stRes.Balance)
	fmt.Println("Transactions:")
	for _, tx := range stRes.Transactions {
		fmt.Printf("(%s) %s -> %s: %d FST\n", tx.Id[:8], tx.Sender[:8], tx.Recipient[:8], tx.Ammount)
	}
}

func (w *WalletNode) balance() {
	blReq := BalanceRequest{WalletAddress: w.WalletAddress}
	req := RPCRequest{Method: "get_bl", Params: blReq}
	rpc := RPCConnection{Type: "tcp", Destination: w.CoreAddress, TimeOut: 3 * time.Second}
	res, err := rpc.Send(req)
	if err != nil {
		fmt.Printf("[ERROR] Network Connection Error: wallet-%v\n", err)
		return
	}
	if !res.Success {
		fmt.Printf("[ERROR] Balance Fetching Error: core-%s\n", res.Error)
		return
	}
	bytes, _ := json.Marshal(res.Data)
	var blRes BalanceResponse
	json.Unmarshal(bytes, &blRes)
	fmt.Println("[SUCCESS] Balance Fetching Successful")
	fmt.Printf("Wallet:\n%s\nBalance:\n%d FST\n", w.WalletAddress, blRes.Balance)
}

func (w *WalletNode) help() {
	fmt.Println("      Command      |            Structure            ")
	fmt.Println("Send Currency      |  tx <ammount> <recipient_pubKey>")
	fmt.Println("Check Statement    |  st                             ")
	fmt.Println("Check Balance      |  bl                             ")
	fmt.Println("Command List Help  |  h                              ")
	fmt.Println("Quit Program       |  q                              ")
}
