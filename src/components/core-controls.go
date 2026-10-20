package components

import "fmt"

func (c *CoreNode) control(usrInput []string) bool {
	switch usrInput[0] {
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

func (c *CoreNode) help() {
	fmt.Println("      Command      |            Structure            ")
	fmt.Println("Command List Help  |  h                              ")
	fmt.Println("Quit Program       |  q                              ")
}
