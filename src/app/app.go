package app

import (
	"bufio"
	"crypto/ecdsa"
	"errors"
	"fastcoin/src/components"
	"fastcoin/src/handlers"
	"fastcoin/src/utils"
	"fmt"
	"os"
	"strings"
	"time"
)

func Run(keys string, node string) error {
	var err error
	switch node {
	case "wallet":
		var wallet *components.WalletNode
		fmt.Println("[INFO] Launching wallet...")
		if !utils.FileExists(keys) {
			wallet, err = newWalletNode("127.0.0.1:9000", keys)
			if err != nil {
				break
			}
		} else {
			wallet, err = loadWalletNode("127.0.0.1:9000", keys)
			if err != nil {
				break
			}
		}
		err = wallet.Enter(keys)
	case "core":
		var core *components.CoreNode
		fmt.Println("[INFO] Launching core...")
		if !utils.FileExists(keys) {
			err = errors.New("missing miner wallet keys")
			fmt.Printf("[FATAL] Keys Inexistence Error: core-%v\n", err)
			break
		}
		core, err = newCoreNode("127.0.0.1:9000", keys)
		if err != nil {
			break
		}
		err = core.Enter(keys)
	default:
		fmt.Println("[FATAL] Unknown Node Error: appinit-invalid node type")
		err = errors.New("invalid node type")
	}
	return err
}

func newWalletNode(coreAddress string, keys string) (*components.WalletNode, error) {
	var err error
	var privKey *ecdsa.PrivateKey
	var pubKey *ecdsa.PublicKey
	fmt.Println("[INFO] No keys found, generate new ones? [y/n]")
	fmt.Printf("fastcoin-wallet@root> ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err = scanner.Err(); err != nil {
			fmt.Printf("[FATAL] Input Read Error: wallet-%v\n", err)
			return nil, err
		}
		fmt.Println("[INFO] Closing wallet...")
		os.Exit(0)
	}
	usrInput := strings.Fields(scanner.Text())
	if !strings.EqualFold(usrInput[0], "y") {
		fmt.Println("[INFO] Closing wallet...")
		os.Exit(0)
	}
	privKey, pubKey, err = handlers.KeysGenerate()
	if err != nil {
		fmt.Printf("[FATAL] Key Generation Error: wallet-%v\n", err)
		return nil, err
	}
	err = handlers.KeysStore(keys, privKey, pubKey)
	if err != nil {
		fmt.Printf("[FATAL] Key Storage Error: wallet-%v\n", err)
		return nil, err
	}
	walletAddress, err := handlers.AddressGenerate(pubKey)
	if err != nil {
		fmt.Printf("[FATAL] Address Generation Error: wallet-%v\n", err)
		return nil, err
	}
	nonce := time.Now().UnixNano()
	fmt.Println("[INFO] Welcome to your new FastCoin wallet, your wallet was successfully created and your keys are stored in keys.pem.")
	return &components.WalletNode{
		PrivKey:       privKey,
		PubKey:        pubKey,
		WalletAddress: walletAddress,
		CoreAddress:   coreAddress,
		Nonce:         nonce,
	}, err
}

func loadWalletNode(coreAddress string, keys string) (*components.WalletNode, error) {
	var err error
	var privKey *ecdsa.PrivateKey
	var pubKey *ecdsa.PublicKey
	fmt.Println("[INFO] Keys found, loading wallet...")
	privKey, pubKey, err = handlers.KeysLoad(keys)
	if err != nil {
		fmt.Printf("[FATAL] Key Load Error: wallet-%v\n", err)
		return nil, err
	}
	walletAddress, err := handlers.AddressGenerate(pubKey)
	if err != nil {
		fmt.Printf("[FATAL] Address Generation Error: wallet-%v\n", err)
		return nil, err
	}
	nonce := time.Now().UnixNano()
	fmt.Println("[INFO] Welcome back to your wallet, type a command and it's parameters (type h for help).")
	return &components.WalletNode{
		PrivKey:       privKey,
		PubKey:        pubKey,
		WalletAddress: walletAddress,
		CoreAddress:   coreAddress,
		Nonce:         nonce,
	}, err
}

func newCoreNode(coreAddress string, keys string) (*components.CoreNode, error) {
	var err error
	var pubKey *ecdsa.PublicKey
	fmt.Println("[INFO] Keys found, creating core...")
	_, pubKey, err = handlers.KeysLoad(keys)
	if err != nil {
		fmt.Printf("[FATAL] Key Load Error: core-%v\n", err)
		return nil, err
	}
	walletAddress, err := handlers.AddressGenerate(pubKey)
	if err != nil {
		fmt.Printf("[FATAL] Address Generation Error: core-%v\n", err)
		return nil, err
	}
	//blockchain handling
	peers := make([]string, 0)
	listenErrChan := make(chan error)
	fmt.Println("[INFO] Welcome to core node, type a command and it's parameters (type h for help).")
	return &components.CoreNode{
		Blockchain:/*blockchain handlers*/ nil,
		Peers:         peers,
		CoreAddress:   coreAddress,
		WalletAddress: walletAddress,
		ListenErrChan: listenErrChan,
		Listening:     false,
	}, err
}
