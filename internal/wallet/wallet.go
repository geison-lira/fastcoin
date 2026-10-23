package wallet

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (w *WalletNode) Enter(keys string) error {
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("fastcoin-wallet@%s...> ", w.WalletAddress[:8])
		if !scanner.Scan() {
			err = scanner.Err()
			if err != nil {
				fmt.Printf("[FATAL] Input Read Error: wallet-%v\n", err)
				return err
			}
			fmt.Println("[INFO] Closing wallet...")
			return err
		}
		usrInput := strings.Fields(scanner.Text())
		if len(usrInput) == 0 {
			continue
		}
		exit := w.control(usrInput)
		if exit {
			return err
		}
	}
}
