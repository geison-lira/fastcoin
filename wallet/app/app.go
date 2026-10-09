package app

import (
	"time"
	"wallet/modules"
)

func Run(keysPath string) {
	wallet := &modules.WalletApp{CoreAddress: "127.0.0.1:9000", Nonce: time.Now().UnixNano()}
	wallet.Enter(keysPath)
}
