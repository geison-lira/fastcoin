package main

import (
	"fastcoin/src/app"
	"flag"
	"fmt"
)

func main() {
	node := flag.String("node", "wallet", "FastCoin node type choice.")
	keys := flag.String("keys", "keys.pem", "public and private keys file location.")
	flag.Parse()
	switch *node {
	case "wallet", "w":
		fmt.Println("Wallet mode.")
	case "full", "f":
		fmt.Println("Full mode.")
	default:
		fmt.Println("Unknown mode.")
	}

	app.Run(*keys)
}
