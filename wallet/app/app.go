package app

import (
	"crypto/ecdsa"
	"fmt"
	"strings"
	"wallet/utils"
)

var privKey *ecdsa.PrivateKey
var pubKey *ecdsa.PublicKey

func Run(keysPath string) {
	fmt.Println("- Launching wallet...")
	if !utils.FileExists(keysPath) {
		runWalletCreationWorkflow(keysPath)
	} else {
		runWalletExecutionWorkflow(keysPath)
	}
	for {
		var usrInput string
		fmt.Printf("> ")
		fmt.Scanln(&usrInput)
		usrInputSliced := strings.Fields(usrInput)
		switch usrInputSliced[0] {
		case "tx":
			runSendCommandWorkflow()
		case "st":
			runStatementCommandWorkflow()
		case "bl":
			runBalanceCommandWorkflow()
		case "h":
			runHelpCommandWorkflow()
		default:
			runUnknownCommandWorkflow()
		}
	}
}
