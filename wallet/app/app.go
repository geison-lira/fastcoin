package app

import (
	"time"
	"wallet/components"
)

func Run(keysPath string) {
	wallet := &components.WalletApp{CoreAddress: "127.0.0.1:9000", Nonce: time.Now().UnixNano()}
	wallet.Enter(keysPath)
}
