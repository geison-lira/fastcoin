package core

import (
	"fastcoin/internal/rpc"
	"fmt"
	"time"
)

func (c *CoreNode) control(usrInput []string) bool {
	switch usrInput[0] {
	case "lt":
		if len(usrInput) == 2 {
			c.listen(usrInput[1:])
		} else {
			fmt.Println("[CORE-ERROR] Command Typing Error: core-wrong command syntax (type h for help).")
		}
	case "mn":
		if len(usrInput) == 2 {
			c.mine(usrInput[1:])
		} else {
			fmt.Println("[CORE-ERROR] Command Typing Error: core-wrong command syntax (type h for help).")
		}
	case "h":
		if len(usrInput) == 1 {
			c.help()
		} else {
			fmt.Println("[CORE-ERROR] Command Typing Error: core-wrong command syntax (type h for help).")
		}
	case "q":
		if len(usrInput) == 1 {
			fmt.Println("[CORE-INFO] Exiting core...")
			return true
		} else {
			fmt.Println("[CORE-ERROR] Command Typing Error: core-wrong command syntax (type h for help).")
		}
	default:
		fmt.Println("[CORE-ERROR] Command Typing Error: core-unknown command (type h for help).")
	}
	return false
}

func (c *CoreNode) listen(usrInput []string) {
	switch usrInput[0] {
	case "start":
		if c.Listening {
			fmt.Println("[CORE-ERROR] Network Listening Error: core-already listening network")
			return
		}
		var rpcConn rpc.RPCConnection
		rpcConn.NewConnection("tcp", c.CoreAddress, 3*time.Second)
		err := rpcConn.Handle(c.ListenSigStop)
		if err != nil {
			fmt.Printf("[CORE-ERROR] Network Connection Error: core-%v\n", err)
			return
		}
		c.Listening = true
		fmt.Println("[CORE-SUCCESS] Core Successfully Listening")
	case "check":
		if !c.Listening {
			fmt.Println("[CORE-INFO] Core is currently not listening.")
			return
		}
		fmt.Println("[CORE-INFO] Core is currently listening")
	case "stop":
		if !c.Listening {
			fmt.Println("[CORE-ERROR] Network Listening Error: core-already not listening network")
			return
		}
		c.ListenSigStop <- true
		c.Listening = false
		fmt.Println("[CORE-SUCCESS] Core Succesfully Stopped Listening")
	default:
		fmt.Println("[CORE-ERROR] Command Parameter Error: core-unknown command parameter (type h for help).")
	}
}

func (c *CoreNode) mine(usrInput []string) {
	switch usrInput[0] {
	case "start":
		if c.Mining {
			fmt.Println("[CORE-ERROR] Blockchain Mining Error: core-already mining blocks")
			return
		}
		//Mining startup here
		c.Mining = true
		fmt.Println("[CORE-SUCCESS] Core Successfully Mining")
	case "check":
		if !c.Listening {
			fmt.Println("[CORE-INFO] Core is currently not listening.")
			return
		}
		fmt.Println("[CORE-INFO] Core is currently listening")
	case "stop":
		if !c.Mining {
			fmt.Println("[CORE-ERROR] Blockchain Mining Error: core-already not mining blocks")
			return
		}
		//Mining stop here
		c.Listening = false
		fmt.Println("[CORE-SUCCESS] Core Succesfully Stopped Mining")
	default:
		fmt.Println("[CORE-ERROR] Command Parameter Error: core-unknown command parameter (type h for help).")
	}
}

func (c *CoreNode) help() {
	fmt.Println("      Command      |            Structure            ")
	fmt.Println("Listening Control  |  lt <mode: start | check | stop>")
	fmt.Println("Mining Control     |  mn <mode: start | check | stop>")
	fmt.Println("Command List Help  |  h                              ")
	fmt.Println("Quit Program       |  q                              ")
}
