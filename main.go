package main

import (
	"fastcoin/internal/app"
	"flag"
	"os"
)

func main() {
	node := flag.String("node", "wallet", "core or wallet node choice.")
	keys := flag.String("keys", "keys.pem", "public and private keys file path.")
	flag.Parse()
	err := app.Run(*node, *keys)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
