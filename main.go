package main

import "fastcoin/src/app"

const keysPath = "keys.pem"

func main() {
	app.Run(keysPath)
}
