package components

import (
	"fmt"
	"time"
)

func (c *CoreNode) control(usrInput []string) bool {
	switch usrInput[0] {
	case "lt":
		if len(usrInput) == 2 {
			c.listen(usrInput[1:])
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

func (c *CoreNode) listen(usrInput []string) {
	switch usrInput[0] {
	case "start":
		if c.Listening {
			fmt.Println("[ERROR] Network Connection Error: core-already listening network")
			return
		}
		var rpc RPCConnection
		rpc.NewConnection("tcp", c.CoreAddress, 3*time.Second)
		err := rpc.Handle(c.ListenErrChan, c.ListenSigStop)
		if err != nil {
			fmt.Printf("[ERROR] Network Connection Error: core-%v\n", err)
			return
		}
		c.Listening = true
		fmt.Println("[SUCCESS] Core Successfully Listening")
	case "check":
		err := <-c.ListenErrChan
		if err != nil {
			fmt.Printf("[ERROR] Network Connection Error: core-%v\n", err)
			return
		}
		fmt.Println("[SUCCESS] Core Successfully Listening")
	case "stop":
		if !c.Listening {
			fmt.Println("[ERROR] Network Connection Error: core-already not listening network")
			return
		}
		c.ListenSigStop <- true
		c.Listening = false
		fmt.Println("[SUCCESS] Core Stopped Listening")
	}
}

func (c *CoreNode) help() {
	fmt.Println("      Command      |            Structure            ")
	fmt.Println("Start Listening    |  lt <mode: start, check, stop>  ")
	fmt.Println("Command List Help  |  h                              ")
	fmt.Println("Quit Program       |  q                              ")
}
