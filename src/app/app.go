package app

import (
	"errors"
	"fastcoin/src/components"
	"fmt"
	"time"
)

func Run(keys string, node string) error {
	var err error
	switch node {
	case "wallet", "w":
		wallet := &components.WalletNode{CoreAddress: "127.0.0.1:9000", Nonce: time.Now().UnixNano()}
		err = wallet.Enter(keys)
	case "core", "c":
		fmt.Println("Core node.")
	default:
		fmt.Println("[FATAL] Unknown Node Error: appinit-invalid node type")
		err = errors.New("invalid node type")
	}
	return err
}
