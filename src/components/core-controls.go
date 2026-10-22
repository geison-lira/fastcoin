package components

import (
	"fmt"
	"time"
)

func (c *CoreNode) control(usrInput []string) bool {
	switch usrInput[0] {
	case "lt":
		if len(usrInput) == 1 {
			c.listen()
		} else {
			fmt.Println("[ERROR] Command Typing Error: core-wrong command syntax (type h for help).")
		}
	case "status":
		if len(usrInput) == 1 {
			c.status()
		} else {
			fmt.Println("[ERROR] Command Typing Error: core-wrong command syntax (type h for help).")
		}
	case "mn":
		fmt.Println("Mining sequence")
	case "h":
		if len(usrInput) == 1 {
			c.help()
		} else {
			fmt.Println("[ERROR] Command Typing Error: core-wrong command syntax (type h for help).")
		}
	case "q":
		if len(usrInput) == 1 {
			fmt.Println("[INFO] Exiting core...")
			return true
		} else {
			fmt.Println("[ERROR] Command Typing Error: core-wrong command syntax (type h for help).")
		}
	default:
		fmt.Println("[ERROR] Command Typing Error: core-unknown command (type h for help).")
	}
	return false
}

func (c *CoreNode) listen() {
	if c.Listening {
		fmt.Println("[ERROR] Network Connection Error: core-already listening network")
		return
	}
	var rpc RPCConnection
	rpc.NewConnection("tcp", c.CoreAddress, 3*time.Second)
	err := rpc.Handle()
	if err != nil {
		fmt.Printf("[ERROR] Network Connection Error: core-%v\n", err)
		return
	}
	c.Listening = true
}

func (c *CoreNode) status() {
	//err := <-goroutine_channel
	if err != nil {
		fmt.Printf("[ERROR] Network Connection Error: core-%v\n", err)
	} else {
		fmt.Println("[SUCCESS] Core Successfully Listening")
	}
}

func (c *CoreNode) help() {
	fmt.Println("      Command      |            Structure            ")
	fmt.Println("Listening Status   |  status                         ")
	fmt.Println("Command List Help  |  h                              ")
	fmt.Println("Quit Program       |  q                              ")
}
