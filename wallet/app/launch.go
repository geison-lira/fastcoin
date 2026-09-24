package app

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"strings"
	"wallet/handlers"
	"wallet/utils"
)

var privKey *ecdsa.PrivateKey
var pubKey *ecdsa.PublicKey

func Run(keysPath string) {
	fmt.Println("- Launching wallet...")
	if !utils.FileExists(keysPath) {
		var usrInput string
		fmt.Println("> No keys found, generate new ones? [y/n]: ")
		fmt.Scanln(&usrInput)
		if strings.EqualFold(usrInput, "y") {
			var err error
			privKey, pubKey, err = handlers.KeysGenerate(keysPath)
			if err != nil {
				fmt.Printf("# Error generating keys: %v\n", err)
				os.Exit(1)
			}
			err = handlers.KeysStore(keysPath, privKey, pubKey)
			if err != nil {
				fmt.Printf("# Error storing keys: %v\n", err)
				os.Exit(1)
			}
			runWalletCreationWorkflow()
		} else {
			fmt.Println("- Closing wallet...")
			os.Exit(0)
		}
	} else {
		fmt.Println("- Keys found, loading wallet...")
		var err error
		privKey, pubKey, err = handlers.KeysLoad(keysPath)
		if err != nil {
			fmt.Printf("# Error loading keys: %v\n", err)
			os.Exit(1)
		}
		runWalletExecutionWorkflow()
	}
	//for loop of transactions
}
