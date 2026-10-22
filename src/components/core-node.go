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
	Listening       bool
	mu              sync.Mutex
}

func (c *CoreNode) Enter(keys string) error {
	var err error
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
