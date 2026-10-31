package app

import (
	"errors"
	"fastcoin/internal/core"
	"fastcoin/internal/utils"
	"fastcoin/internal/wallet"
	"fmt"
)

func Run(node string, keys string, peers string) error {
	var err error
	switch node {
	case "wallet":
		var wallet wallet.WalletNode
		fmt.Println("[INIT-INFO] Launching wallet...")
		if !utils.FileExists(peers) {
			err = errors.New("missing peer nodes file")
			fmt.Printf("[INIT-FATAL] Peers Inexistence Error: init-%v\n", err)
			break
		}
		if !utils.FileExists(keys) {
			err = wallet.NewWalletNode(peers, keys)
			if err != nil {
				break
			}
		} else {
			err = wallet.LoadWalletNode(peers, keys)
			if err != nil {
				break
			}
		}
		err = wallet.Enter(keys)
	case "core":
		var core core.CoreNode
		fmt.Println("[INIT-INFO] Launching core...")
		if !utils.FileExists(keys) {
			err = errors.New("missing miner wallet keys")
			fmt.Printf("[INIT-FATAL] Keys Inexistence Error: init-%v\n", err)
			break
		}
		err = core.NewCoreNode("127.0.0.1:9000", keys)
		if err != nil {
			break
		}
		err = core.Enter(keys)
	default:
		err = errors.New("invalid node type")
		fmt.Printf("[INIT-FATAL] Unknown Node Error: init-%v\n", err)
	}
	return err
}
