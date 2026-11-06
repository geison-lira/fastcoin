package app

import (
	"errors"
	"fastcoin/internal/core"
	"fastcoin/internal/wallet"
	"fmt"
)

func Run(node string, keys string, peers string) error {
	var err error
	switch node {
	case "wallet":
		var wallet wallet.WalletNode
		fmt.Println("[INIT-INFO] Launching wallet...")
		err = wallet.NewWalletNode(peers, keys)
		if err != nil {
			return err
		}
		err = wallet.Enter()
		if err != nil {
			return err
		}
	case "core":
		var core core.CoreNode
		fmt.Println("[INIT-INFO] Launching core...")
		err = core.NewCoreNode("127.0.0.1:9000", keys)
		if err != nil {
			return err
		}
		err = core.Enter()
		if err != nil {
			return err
		}
	default:
		err = errors.New("invalid node type")
		fmt.Printf("[INIT-FATAL] Unknown Node Error: init-%v\n", err)
		return err
	}
	return nil
}
