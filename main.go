package main

import (
	"fastcoin/src/app"
	"flag"
	"os"
)

func main() {
	node := flag.String("node", "wallet", "full or wallet node choice.")
	keys := flag.String("keys", "keys.pem", "public and private keys file path.")
	flag.Parse()
	err := app.Run(*keys, *node)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
