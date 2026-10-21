package components

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func (w *WalletNode) control(usrInput []string) bool {
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
			return true
		} else {
			fmt.Println("[ERROR] Command Typing Error: wallet-wrong command syntax (type h for help).")
		}
	default:
		fmt.Println("[ERROR] Command Typing Error: wallet-unknown command (type h for help).")
	}
	return false
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
	var tx *Transaction
	tx.NewTransaction(w.WalletAddress, recipient, ammount, w.Nonce)
	signBytes, err := ecdsa.SignASN1(rand.Reader, w.PrivKey, []byte(tx.Id))
	if err != nil {
		fmt.Printf("[ERROR] Transaction Signing Error: wallet-%v\n", err)
		return
	}
	tx.Sign(signBytes)
	var req RPCRequest
	req.NewRequest("set_tx", tx)
	var rpc RPCConnection
	rpc.NewConnection("tcp", w.CoreAddress, 3*time.Second)
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
	var stReq *StatementRequest
	stReq.NewRequest(w.WalletAddress)
	var req RPCRequest
	req.NewRequest("get_st", stReq)
	var rpc RPCConnection
	rpc.NewConnection("tcp", w.CoreAddress, 3*time.Second)
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
	fmt.Printf("Wallet:\n%s\nBalance:\n%d FST\n", stRes.WalletAddress, stRes.Balance)
	fmt.Println("Transactions:")
	for _, tx := range stRes.Transactions {
		fmt.Printf("(%s) %s -> %s: %d FST\n", tx.Id[:8], tx.Sender[:8], tx.Recipient[:8], tx.Ammount)
	}
}

func (w *WalletNode) balance() {
	var blReq *BalanceRequest
	blReq.NewRequest(w.WalletAddress)
	var req RPCRequest
	req.NewRequest("get_bl", blReq)
	var rpc RPCConnection
	rpc.NewConnection("tcp", w.CoreAddress, 3*time.Second)
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
	fmt.Printf("Wallet:\n%s\nBalance:\n%d FST\n", blRes.WalletAddress, blRes.Balance)
}

func (w *WalletNode) help() {
	fmt.Println("      Command      |            Structure            ")
	fmt.Println("Send Currency      |  tx <ammount> <recipient_pubKey>")
	fmt.Println("Check Statement    |  st                             ")
	fmt.Println("Check Balance      |  bl                             ")
	fmt.Println("Command List Help  |  h                              ")
	fmt.Println("Quit Program       |  q                              ")
}
