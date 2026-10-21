package components

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

const difficulty = 3

type CoreNode struct {
	Blockchain      *Blockchain
	Peers           []string
	CoreAddress     string
	WalletAddress   string
	ListenerChannel chan error
	mu              sync.Mutex
}

func (c *CoreNode) Enter(keys string) error {
	var err error
	err = c.listen()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("fastcoin-core@%s...> ", c.WalletAddress[:8])
		if !scanner.Scan() {
			err = scanner.Err()
			if err != nil {
				fmt.Printf("[FATAL] Input Read Error: core-%v\n", err)
				return err
			}
			fmt.Println("[INFO] Closing core...")
			os.Exit(0)
		}
		usrInput := strings.Fields(scanner.Text())
		if len(usrInput) == 0 {
			continue
		}
		exit := c.control(usrInput)
		if exit {
			return err
		}
	}
}

func (c *CoreNode) listen() error {
	rpc := RPCConnection{Type: "tcp", Address: c.CoreAddress}
	listener, err := rpc.Listen()
	if err != nil {
		fmt.Printf("[ERROR] Network Connection Error: core-%v\n", err)
		return err
	}
	go rpc.Receive(listener, c.ListenerChannel)
	return nil
}
