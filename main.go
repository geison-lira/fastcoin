package main

import (
	"fastcoin/internal/app"
	"flag"
	"os"
)

func main() {
	node := flag.String("node", "wallet", "core or wallet node choice.")
	keys := flag.String("keys", "keys.pem", "public and private keys file path.")
	peers := flag.String("peers", "peers.txt", "peer nodes addresses file path.")
	flag.Parse()
	err := app.Run(*node, *keys, *peers)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
