package wallet

import (
	"encoding/json"
	"fastcoin/internal/handlers"
	"fastcoin/internal/rpc"
	"fmt"
	"strconv"
)

func (w *WalletNode) control(usrInput []string) bool {
	switch usrInput[0] {
	case "tx":
		if len(usrInput) == 3 {
			w.transaction(usrInput[1:])
		} else {
			fmt.Println("[WALLET-ERROR] Command Typing Error: wallet-wrong command syntax (type h for help).")
		}
	case "st":
		if len(usrInput) == 1 {
			w.statement()
		} else {
			fmt.Println("[WALLET-ERROR] Command Typing Error: wallet-wrong command syntax (type h for help).")
		}
	case "bl":
		if len(usrInput) == 1 {
			w.balance()
		} else {
			fmt.Println("[WALLET-ERROR] Command Typing Error: wallet-wrong command syntax (type h for help).")
		}
	case "h":
		if len(usrInput) == 1 {
			w.help()
		} else {
			fmt.Println("[WALLET-ERROR] Command Typing Error: wallet-wrong command syntax (type h for help).")
		}
	case "q":
		if len(usrInput) == 1 {
			fmt.Println("[WALLET-INFO] Exiting wallet...")
			return true
		} else {
			fmt.Println("[WALLET-ERROR] Command Typing Error: wallet-wrong command syntax (type h for help).")
		}
	default:
		fmt.Println("[WALLET-ERROR] Command Typing Error: wallet-unknown command (type h for help).")
	}
	return false
}

func (w *WalletNode) transaction(usrInput []string) {
	ammount, err := strconv.ParseInt(usrInput[0], 10, 64)
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Transaction Parsing Error: wallet-%v\n", err)
		return
	}
	if ammount <= 0 {
		fmt.Println("[WALLET-ERROR] Transaction Ammount Error: wallet-invalid ammount")
		return
	}
	recipient := usrInput[1]
	w.Nonce++
	var tx rpc.Transaction
	pubKeyString, err := handlers.PubKeyToString(w.PubKey)
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Transaction Parsing Error: wallet-%v\n", err)
		return
	}
	err = tx.NewTransaction(w.WalletAddress, recipient, ammount, w.Nonce, w.PrivKey, pubKeyString)
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Transaction Signing Error: wallet-%v\n", err)
		return
	}
	var req rpc.RPCRequest
	err = req.NewRequest("set_tx", tx)
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Transaction Parsing Error: wallet-%v\n", err)
		return
	}
	res, err := w.Connection.Send(req)
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Network Connection Error: wallet-%v\n", err)
		return
	}
	if !res.Success {
		fmt.Printf("[WALLET-ERROR] Transaction Broadcast Error: core-%s\n", res.Error)
		return
	}
	fmt.Println("[WALLET-SUCCESS] Transaction Broadcast Successful")
	fmt.Printf("TxId: %s\n", tx.Id)
}

func (w *WalletNode) statement() {
	var stReq rpc.StatementRequest
	stReq.NewRequest(w.WalletAddress)
	var req rpc.RPCRequest
	err := req.NewRequest("get_st", stReq)
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Statement Parsing Error: wallet-%v\n", err)
		return
	}
	res, err := w.Connection.Send(req)
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Network Connection Error: wallet-%v\n", err)
		return
	}
	if !res.Success {
		fmt.Printf("[WALLET-ERROR] Statement Fetching Error: core-%s\n", res.Error)
		return
	}
	var stRes rpc.StatementResponse
	json.Unmarshal(res.Data, &stRes)
	fmt.Println("[WALLET-SUCCESS] Statement Fetching Successful")
	fmt.Printf("Wallet:\n%s\nBalance:\n%d FST\n", stRes.WalletAddress, stRes.Balance)
	fmt.Println("Transactions:")
	for _, tx := range stRes.Transactions {
		fmt.Printf("(%s) %s -> %s: %d FST\n", tx.Id[:8], tx.Sender[:8], tx.Recipient[:8], tx.Ammount)
	}
}

func (w *WalletNode) balance() {
	var blReq rpc.BalanceRequest
	blReq.NewRequest(w.WalletAddress)
	var req rpc.RPCRequest
	err := req.NewRequest("get_bl", blReq)
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Balance Parsing Error: wallet-%v\n", err)
		return
	}
	res, err := w.Connection.Send(req)
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Network Connection Error: wallet-%v\n", err)
		return
	}
	if !res.Success {
		fmt.Printf("[WALLET-ERROR] Balance Fetching Error: core-%s\n", res.Error)
		return
	}
	var blRes rpc.BalanceResponse
	json.Unmarshal(res.Data, &blRes)
	fmt.Println("[WALLET-SUCCESS] Balance Fetching Successful")
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
