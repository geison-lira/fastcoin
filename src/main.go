package main

import "src/app"

const keysPath = "keys.pem"

func main() {
	app.Run(keysPath)
}
