package app

import "fmt"

func runWalletCreationWorkflow() {
	fmt.Println("[NEW WALLET CREATION]")
	fmt.Println("Welcome to the Minicoin Blockchain, your wallet was successfully created and your keys are stored in keys.pem.")
}

func runWalletExecutionWorkflow() {
	fmt.Println("[WALLET EXECUTION]")
	fmt.Println("Welcome back to your Minicoin Wallet, ready to process transactions.")
}
