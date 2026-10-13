package app

import (
	"errors"
	"fastcoin/src/components"
	"fmt"
	"time"
)

func Run(keys string, node string) error {
	switch node {
	case "wallet", "w":
		wallet := &components.WalletApp{CoreAddress: "127.0.0.1:9000", Nonce: time.Now().UnixNano()}
		wallet.Enter(keys)
	case "full", "f":
		fmt.Println("Full mode.")
	default:
		fmt.Println("[FATAL] Unknown Node Error: invalid node type")
		return errors.New("invalid node type")
	}
	return nil
}
