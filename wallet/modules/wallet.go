package modules

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"os"
	"strings"
	"wallet/handlers"
	"wallet/utils"
)

type WalletApp struct {
	PrivKey     *ecdsa.PrivateKey
	PubKey      *ecdsa.PublicKey
	Address     string
	CoreAddress string
	Nonce       int64
}

func (w *WalletApp) Enter(keysPath string) {
	fmt.Println("sys@root> Launching wallet...")
	if !utils.FileExists(keysPath) {
		w.walletCreate(keysPath)
	} else {
		w.walletLoad(keysPath)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("sys@%s...> ", w.Address[:8])
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("sys@root> Error reading input: %v\n", err)
				os.Exit(1)
			} else {
				fmt.Println("\nsys@root> Closing wallet...")
				os.Exit(0)
			}
		}
		usrInputSlice := strings.Fields(scanner.Text())
		if len(usrInputSlice) == 0 {
			continue
		}
		switch usrInputSlice[0] {
		case "tx":
			if len(usrInputSlice) == 3 {
				w.transactionSend(usrInputSlice[1:])
			} else {
				w.commandError()
			}
		case "st":
			if len(usrInputSlice) == 1 {
				w.statementCheck()
			} else {
				w.commandError()
			}
		case "bl":
			if len(usrInputSlice) == 1 {
				w.balanceCheck()
			} else {
				w.commandError()
			}
		case "h":
			if len(usrInputSlice) == 1 {
				w.helpView()
			} else {
				w.commandError()
			}
		default:
			w.unknownError()
		}
	}
}

func (w *WalletApp) walletCreate(keysPath string) {
	fmt.Printf("sys@root> No keys found, generate new ones? [y/n]: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("sys@root> Error reading input: %v\n", err)
			os.Exit(1)
		} else {
			fmt.Println("\nsys@root> Closing wallet...")
			os.Exit(0)
		}
	}
	usrInputSlice := strings.Fields(scanner.Text())
	if strings.EqualFold(usrInputSlice[0], "y") {
		var err error
		w.PrivKey, w.PubKey, err = handlers.KeysGenerate()
		if err != nil {
			fmt.Printf("sys@root> Error generating keys: %v\n", err)
			os.Exit(1)
		}
		err = handlers.KeysStore(keysPath, w.PrivKey, w.PubKey)
		if err != nil {
			fmt.Printf("sys@root> Error storing keys: %v\n", err)
			os.Exit(1)
		}
		w.Address, err = handlers.AddressGenerate(w.PubKey)
		if err != nil {
			fmt.Printf("sys@root> Error generating address: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("sys@root> Welcome to the Fastcoin Blockchain, your wallet was successfully created and your keys are stored in keys.pem.")
	} else {
		fmt.Println("\nsys@root> Closing wallet...")
		os.Exit(0)
	}
}

func (w *WalletApp) walletLoad(keysPath string) {
	fmt.Println("sys@root> Keys found, loading wallet...")
	var err error
	w.PrivKey, w.PubKey, err = handlers.KeysLoad(keysPath)
	if err != nil {
		fmt.Printf("sys@root> Error loading keys: %v\n", err)
		os.Exit(1)
	}
	w.Address, err = handlers.AddressGenerate(w.PubKey)
	fmt.Printf("sys@root> Welcome back to your wallet, type a command and it's parameters (type h for help).")
}

func (w *WalletApp) transactionSend(usrInputSlice []string) {
	//
	fmt.Println("- Sending transaction...")
	//
}

func (w *WalletApp) statementCheck() {
	//
	fmt.Println("- Fetching statement...")
	//
}

func (w *WalletApp) balanceCheck() {
	//
	fmt.Println("- Fetching balance...")
	//
}

func (w *WalletApp) helpView() {
	fmt.Println("     Command     |            Structure            ")
	fmt.Println("-Send Currency   | tx <ammount> <recipient_pubKey> ")
	fmt.Println("-Check Statement | st                              ")
	fmt.Println("-Check Balance   | bl                              ")
	fmt.Println("-Help            | h                               ")
}

func (w *WalletApp) commandError() {
	fmt.Println("sys@root> Error in command structure, type h to see all commands.")
}

func (w *WalletApp) unknownError() {
	fmt.Println("sys@root> Error in command, type h to see all commands.")
}
