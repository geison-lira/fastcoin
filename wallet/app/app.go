package app

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		if !scanner.Scan() {
			break
		}
		usrInput := scanner.Text()
		usrInputSliced := strings.Fields(usrInput)
		if len(usrInputSliced) == 0 {
			continue
		}
		switch usrInputSliced[0] {
		case "tx":
			if len(usrInputSliced) == 3 {
				runSendCommandWorkflow()
			} else {
				runWrongCommandWorkflow()
			}
		case "st":
			if len(usrInputSliced) == 1 {
				runStatementCommandWorkflow()
			} else {
				runWrongCommandWorkflow()
			}
		case "bl":
			if len(usrInputSliced) == 1 {
				runBalanceCommandWorkflow()
			} else {
				runWrongCommandWorkflow()
			}
		case "h":
			if len(usrInputSliced) == 1 {
				runHelpCommandWorkflow()
			} else {
				runWrongCommandWorkflow()
			}
		default:
			runUnknownCommandWorkflow()
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("# Error reading input: %v\n", err)
	} else {
		fmt.Println("\n- Closing wallet...")
	}
}
