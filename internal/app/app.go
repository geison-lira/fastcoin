package app

import (
	"errors"
	"fastcoin/internal/core"
	"fastcoin/internal/utils"
	"fastcoin/internal/wallet"
	"fmt"
)

func Run(keys string, node string) error {
	var err error
	switch node {
	case "wallet":
		var wallet *wallet.WalletNode
		fmt.Println("[INFO] Launching wallet...")
		if !utils.FileExists(keys) {
			err = wallet.NewWalletNode("127.0.0.1:9000", keys)
			if err != nil {
				break
			}
		} else {
			err = wallet.LoadWalletNode("127.0.0.1:9000", keys)
			if err != nil {
				break
			}
		}
		err = wallet.Enter(keys)
	case "core":
		var core *core.CoreNode
		fmt.Println("[INFO] Launching core...")
		if !utils.FileExists(keys) {
			err = errors.New("missing miner wallet keys")
			fmt.Printf("[FATAL] Keys Inexistence Error: core-%v\n", err)
			break
		}
		err = core.NewCoreNode("127.0.0.1:9000", keys)
		if err != nil {
			break
		}
		err = core.Enter(keys)
	default:
		fmt.Println("[FATAL] Unknown Node Error: appinit-invalid node type")
		err = errors.New("invalid node type")
	}
	return err
}
