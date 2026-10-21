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
	case "wallet":
		wallet := newWalletNode("127.0.0.1:9000", time.Now().UnixNano())
		err = wallet.Enter(keys)
	case "core":
		core := newCoreNode("127.0.0.1:9000")
		err = core.Enter(keys)
	default:
		fmt.Println("[FATAL] Unknown Node Error: appinit-invalid node type")
		err = errors.New("invalid node type")
	}
	return err
}

func newWalletNode(coreAddress string, nonce int64) *components.WalletNode {
	return &components.WalletNode{CoreAddress: coreAddress, Nonce: nonce}
}

func newCoreNode(coreAddress string) *components.CoreNode {
	return &components.CoreNode{CoreAddress: coreAddress}
}
