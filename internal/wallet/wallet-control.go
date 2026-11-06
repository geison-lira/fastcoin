package wallet

import (
	"encoding/json"
	"errors"
	"fastcoin/internal/handlers"
	"fastcoin/internal/rpc"
	"fmt"
	"strconv"
	"time"
)

func (w *WalletNode) control(usrInput []string) bool {
	var err error
	switch usrInput[0] {
	case "tx":
		if len(usrInput) == 3 {
			w.transaction(usrInput[1:])
		} else {
			err = errors.New("wrong command syntax")
			fmt.Printf("[WALLET-ERROR] Command Typing Error: wallet-%v\n", err)
		}
	case "st":
		if len(usrInput) == 1 {
			w.statement()
		} else {
			err = errors.New("wrong command syntax")
			fmt.Printf("[WALLET-ERROR] Command Typing Error: wallet-%v\n", err)
		}
	case "bl":
		if len(usrInput) == 1 {
			w.balance()
		} else {
			err = errors.New("wrong command syntax")
			fmt.Printf("[WALLET-ERROR] Command Typing Error: wallet-%v\n", err)
		}
	case "ad":
		if len(usrInput) == 1 {
			w.address()
		} else {
			err = errors.New("wrong command syntax")
			fmt.Printf("[WALLET-ERROR] Command Typing Error: wallet-%v\n", err)
		}
	case "h":
		if len(usrInput) == 1 {
			w.help()
		} else {
			err = errors.New("wrong command syntax")
			fmt.Printf("[WALLET-ERROR] Command Typing Error: wallet-%v\n", err)
		}
	case "q":
		if len(usrInput) == 1 {
			return true
		} else {
			err = errors.New("wrong command syntax")
			fmt.Printf("[WALLET-ERROR] Command Typing Error: wallet-%v\n", err)
		}
	default:
		err = errors.New("unknown command")
		fmt.Printf("[WALLET-ERROR] Command Typing Error: wallet-%v\n", err)
	}
	return false
}

func (w *WalletNode) transaction(usrInput []string) {
	var err error
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
	pubKeyString, err := handlers.PubKeyToString(w.PubKey)
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Transaction Parsing Error: wallet-%v\n", err)
		return
	}
	var tx rpc.Transaction
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
	var res *rpc.RPCResponse
	var connection rpc.RPCConnection
	w.Nonce++
	for range 10 {
		coreAddress := handlers.PeersSelect(w.Peers)
		connection.NewConnection("tcp", coreAddress, 3*time.Second)
		res, err = connection.Send(req)
		if err == nil {
			break
		}
	}
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Network Connection Error: wallet-%v\n", err)
		return
	}
	if !res.Success {
		fmt.Printf("[WALLET-ERROR] Transaction Broadcast Error: core-%s\n", res.Error)
		return
	}
	fmt.Println("[WALLET-SUCCESS] Transaction Broadcast Successful")
	fmt.Printf("(%s) %s -> %s: %d FST | Signature: %s\n", tx.Id, tx.Sender, tx.Recipient, tx.Ammount, tx.Signature)
}

func (w *WalletNode) statement() {
	var err error
	var stReq rpc.StatementRequest
	stReq.NewRequest(w.WalletAddress)
	var req rpc.RPCRequest
	err = req.NewRequest("get_st", stReq)
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Statement Parsing Error: wallet-%v\n", err)
		return
	}
	var res *rpc.RPCResponse
	var connection rpc.RPCConnection
	for range 10 {
		coreAddress := handlers.PeersSelect(w.Peers)
		connection.NewConnection("tcp", coreAddress, 3*time.Second)
		res, err = connection.Send(req)
		if err == nil {
			break
		}
	}
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
	fmt.Printf("Wallet: %s\nBalance: %d FST\n", w.WalletAddress, stRes.Balance)
	fmt.Println("Transactions:")
	for _, tx := range stRes.Transactions {
		fmt.Printf("(%s) %s -> %s: %d FST | Signature: %s\n", tx.Id, tx.Sender, tx.Recipient, tx.Ammount, tx.Signature)
	}
}

func (w *WalletNode) balance() {
	var err error
	var blReq rpc.BalanceRequest
	blReq.NewRequest(w.WalletAddress)
	var req rpc.RPCRequest
	err = req.NewRequest("get_bl", blReq)
	if err != nil {
		fmt.Printf("[WALLET-ERROR] Balance Parsing Error: wallet-%v\n", err)
		return
	}
	var res *rpc.RPCResponse
	var connection rpc.RPCConnection
	for range 10 {
		coreAddress := handlers.PeersSelect(w.Peers)
		connection.NewConnection("tcp", coreAddress, 3*time.Second)
		res, err = connection.Send(req)
		if err == nil {
			break
		}
	}
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
	fmt.Printf("Wallet: %s\nBalance: %d FST\n", w.WalletAddress, blRes.Balance)
}

func (w *WalletNode) address() {
	fmt.Println("[WALLET-SUCCESS] Address Fetching Successful")
	fmt.Printf("Wallet: %s\n", w.WalletAddress)
}

func (w *WalletNode) help() {
	fmt.Println("      Command      |            Structure            ")
	fmt.Println("Send Currency      |  tx <ammount> <recipient_pubKey>")
	fmt.Println("Check Statement    |  st                             ")
	fmt.Println("Check Balance      |  bl                             ")
	fmt.Println("Check Address      |  ad                             ")
	fmt.Println("Command List Help  |  h                              ")
	fmt.Println("Quit Program       |  q                              ")
}
