package main

import "wallet/app"

const keysPath = "keys.pem"

func main() {
	app.Run(keysPath)
}
