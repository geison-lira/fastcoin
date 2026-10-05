package app

import "fmt"

func runWalletCreationWorkflow() {
	fmt.Println("[WALLET CREATED]")
	fmt.Println("--#-- Welcome to the Fastcoin Blockchain, your wallet was successfully created and your keys are stored in keys.pem. --#--")
}

func runWalletExecutionWorkflow() {
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

func runUnknownCommandWorkflow() {
	fmt.Println("--#-- Unknown command, type h to see all valid commands. --#--")
}
