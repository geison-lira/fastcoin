package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
