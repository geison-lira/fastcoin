package components

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"wallet/handlers"
	"wallet/utils"
)

type WalletApp struct {
	PrivKey       *ecdsa.PrivateKey
	PubKey        *ecdsa.PublicKey
	WalletAddress string
	CoreAddress   string
	Nonce         int64
}

func (w *WalletApp) Enter(keysPath string) {
	fmt.Println("[INFO] Launching wallet...")
	if !utils.FileExists(keysPath) {
		w.create(keysPath)
	} else {
		w.load(keysPath)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("fastcoin@%s...> ", w.WalletAddress[:8])
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("[FATAL] Input Read Error: %v\n", err)
				os.Exit(1)
			} else {
				fmt.Println("[INFO] Closing wallet...")
				os.Exit(0)
			}
		}
		usrInputSlice := strings.Fields(scanner.Text())
		if len(usrInputSlice) == 0 {
			continue
		}
		switch usrInputSlice[0] {
		case "tx":
			if len(usrInputSlice) == 3 {
				w.transaction(usrInputSlice[1:])
			} else {
				w.commandError()
			}
		case "st":
			if len(usrInputSlice) == 1 {
				w.statement()
			} else {
				w.commandError()
			}
		case "bl":
			if len(usrInputSlice) == 1 {
				w.balance()
			} else {
				w.commandError()
			}
		case "h":
			if len(usrInputSlice) == 1 {
				w.help()
			} else {
				w.commandError()
			}
		default:
			w.unknownError()
		}
	}
}

func (w *WalletApp) create(keysPath string) {
	fmt.Printf("[INFO] No keys found, generate new ones? [y/n]")
	fmt.Printf("fastcoin@root> ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("[FATAL] Input Read Error: %v\n", err)
			os.Exit(1)
		} else {
			fmt.Println("[INFO] Closing wallet...")
			os.Exit(0)
		}
	}
	usrInputSlice := strings.Fields(scanner.Text())
	if strings.EqualFold(usrInputSlice[0], "y") {
		var err error
		w.PrivKey, w.PubKey, err = handlers.KeysGenerate()
		if err != nil {
			fmt.Printf("[FATAL] Key Generation Error: %v\n", err)
			os.Exit(1)
		}
		err = handlers.KeysStore(keysPath, w.PrivKey, w.PubKey)
		if err != nil {
			fmt.Printf("[FATAL] Key Storage Error: %v\n", err)
			os.Exit(1)
		}
		w.WalletAddress, err = handlers.AddressGenerate(w.PubKey)
		if err != nil {
			fmt.Printf("[FATAL] Address Generation Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("[INFO] Welcome to your new FastCoin wallet, your wallet was successfully created and your keys are stored in keys.pem.")
	} else {
		fmt.Println("[INFO] Closing wallet...")
		os.Exit(0)
	}
}

func (w *WalletApp) load(keysPath string) {
	fmt.Println("[INFO] Keys found, loading wallet...")
	var err error
	w.PrivKey, w.PubKey, err = handlers.KeysLoad(keysPath)
	if err != nil {
		fmt.Printf("[FATAL] Key Load Error %v\n", err)
		os.Exit(1)
	}
	w.WalletAddress, err = handlers.AddressGenerate(w.PubKey)
	fmt.Println("[INFO] Welcome back to your wallet, type a command and it's parameters (type h for help).")
}

func (w *WalletApp) transaction(usrInputSlice []string) {
	ammount, err := strconv.ParseInt(usrInputSlice[0], 10, 64)
	if err != nil {
		fmt.Printf("[ERROR] Transaction Parsing Error: %v\n", err)
		return
	}
	if ammount <= 0 {
		fmt.Println("[ERROR] Transaction Ammount Error: invalid ammount")
		return
	}
	recipient := usrInputSlice[1]
	w.Nonce++
	tx := Transaction{Sender: w.WalletAddress, Recipient: recipient, Ammount: ammount, Nonce: w.Nonce}
	tx.Id = tx.Hash()
	signBytes, err := ecdsa.SignASN1(rand.Reader, w.PrivKey, []byte(tx.Id))
	if err != nil {
		fmt.Printf("[ERROR] Transaction Signing Error: %v\n", err)
		return
	}
	tx.Signature = hex.EncodeToString(signBytes)
	req := RPCRequest{Method: "set_tx", Params: tx}
	rpc := RPCConnection{Type: "tcp", Destination: w.CoreAddress, TimeOut: 3 * time.Second}
	res, err := rpc.Send(req)
	if err != nil {
		fmt.Printf("[ERROR] Network Connection Error: %v\n", err)
		return
	}
	if res.Success {
		fmt.Println("[SUCCESS] Transaction Broadcast Successful")
		fmt.Printf("TxId: %s\n", tx.Id)
	} else {
		fmt.Printf("[ERROR] Transaction Broadcast Error: core-%s\n", res.Error)
	}
}

func (w *WalletApp) statement() {
	sReq := StatementRequest{WalletAddress: w.WalletAddress}
	req := RPCRequest{Method: "get_st", Params: sReq}
	rpc := RPCConnection{Type: "tcp", Destination: w.CoreAddress, TimeOut: 3 * time.Second}
	res, err := rpc.Send(req)
	if err != nil {
		fmt.Printf("[ERROR] Network Connection Error: %v\n", err)
		return
	}
	if res.Success {
		bytes, _ := json.Marshal(res.Data)
		var sRes StatementResponse
		json.Unmarshal(bytes, &sRes)
		fmt.Println("[SUCCESS] Statement Fetching Successful")
		fmt.Printf("Wallet:\n%s\nBalance\n%d FST\n", w.WalletAddress, sRes.Balance)
		fmt.Println("Transactions:")
		for _, tx := range sRes.Transactions {
			fmt.Printf("(%s) %s -> %s: %d FST\n", tx.Id[:8], tx.Sender[:8], tx.Recipient[:8], tx.Ammount)
		}
	} else {
		fmt.Printf("[ERROR] Statement Fetching Error: core-%s\n", res.Error)
	}
}

func (w *WalletApp) balance() {
	bReq := BalanceRequest{WalletAddress: w.WalletAddress}
	req := RPCRequest{Method: "get_bl", Params: bReq}
	rpc := RPCConnection{Type: "tcp", Destination: w.CoreAddress, TimeOut: 3 * time.Second}
	res, err := rpc.Send(req)
	if err != nil {
		fmt.Printf("[ERROR] Network Connection Error: %v\n", err)
		return
	}
	if res.Success {
		bytes, _ := json.Marshal(res.Data)
		var bRes BalanceResponse
		json.Unmarshal(bytes, &bRes)
		fmt.Println("[SUCCESS] Balance Fetching Successful")
		fmt.Printf("Wallet:\n%s\nBalance\n%d FST\n", w.WalletAddress, bRes.Balance)
	} else {
		fmt.Printf("[ERROR] Balance Fetching Error: core-%s\n", res.Error)
	}
}

func (w *WalletApp) help() {
	fmt.Println("     Command     |            Structure            ")
	fmt.Println("Send Currency    |  tx <ammount> <recipient_pubKey>")
	fmt.Println("Check Statement  |  st                             ")
	fmt.Println("Check Balance    |  bl                             ")
	fmt.Println("Help             |  h                              ")
}

func (w *WalletApp) commandError() {
	fmt.Println("sys@root> Error in command structure, type h to see all commands.")
}

func (w *WalletApp) unknownError() {
	fmt.Println("sys@root> Error in command, type h to see all commands.")
}
