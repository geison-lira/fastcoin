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
			return err
		}
		if !utils.FileExists(keys) {
			err = wallet.NewWalletNode(peers, keys)
			if err != nil {
				return err
			}
		} else {
			err = wallet.LoadWalletNode(peers, keys)
			if err != nil {
				return err
			}
		}
		err = wallet.Enter()
		if err != nil {
			return err
		}
	case "core":
		var core core.CoreNode
		fmt.Println("[INIT-INFO] Launching core...")
		if !utils.FileExists(keys) {
			err = errors.New("missing miner wallet keys")
			fmt.Printf("[INIT-FATAL] Keys Inexistence Error: init-%v\n", err)
			return err
		}
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
