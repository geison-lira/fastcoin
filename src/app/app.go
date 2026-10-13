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
		wallet := &components.WalletApp{CoreAddress: "127.0.0.1:9000", Nonce: time.Now().UnixNano()}
		err = wallet.Enter(keys)
	case "full", "f":
		fmt.Println("Full mode.")
	default:
		fmt.Println("[FATAL] Unknown Node Error: invalid node type")
		err = errors.New("invalid node type")
	}
	return err
}
