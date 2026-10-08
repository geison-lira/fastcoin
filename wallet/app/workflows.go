package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"wallet/handlers"
)

func runWalletCreationWorkflow(keysPath string) {
	fmt.Printf("> No keys found, generate new ones? [y/n]: ")
	scanner := bufio.NewScanner((os.Stdin))
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("# Error reading input: %v\n", err)
			os.Exit(1)
		} else {
			fmt.Println("\n- Closing wallet...")
			os.Exit(0)
		}
	}
	usrInput := scanner.Text()
	usrInputSliced := strings.Fields(usrInput)
	if strings.EqualFold(usrInputSliced[0], "y") {
		var err error
		privKey, pubKey, err = handlers.KeysGenerate()
		if err != nil {
			fmt.Printf("# Error generating keys: %v\n", err)
			os.Exit(1)
		}
		err = handlers.KeysStore(keysPath, privKey, pubKey)
		if err != nil {
			fmt.Printf("# Error storing keys: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("[WALLET CREATED]")
		fmt.Println("--#-- Welcome to the Fastcoin Blockchain, your wallet was successfully created and your keys are stored in keys.pem. --#--")
	} else {
		fmt.Println("- Closing wallet...")
		os.Exit(0)
	}
}

func runWalletExecutionWorkflow(keysPath string) {
	fmt.Println("- Keys found, loading wallet...")
	var err error
	privKey, pubKey, err = handlers.KeysLoad(keysPath)
	if err != nil {
		fmt.Printf("# Error loading keys: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("[WALLET ACCESSED]")
	fmt.Println("--#-- Welcome back to your wallet, type a command and it's parameters (type h for help). --#--")
}

func runSendCommandWorkflow() {
	fmt.Println("- Sending transaction...")
}

func runStatementCommandWorkflow() {
	fmt.Println("- Fetching statement...")
}

func runBalanceCommandWorkflow() {
	fmt.Println("- Fetching balance...")
}

func runHelpCommandWorkflow() {
	fmt.Println("     Command     |            Structure            ")
	fmt.Println("-Send Currency   | tx <ammount> <recipient_pubKey> ")
	fmt.Println("-Check Statement | st                              ")
	fmt.Println("-Check Balance   | bl                              ")
	fmt.Println("-Help            | h                               ")
}

func runWrongCommandWorkflow() {
	fmt.Println("# Error in command structure, type h to see all commands structures.")
}

func runUnknownCommandWorkflow() {
	fmt.Println("# Error in command, type h to see all valid commands.")
}
