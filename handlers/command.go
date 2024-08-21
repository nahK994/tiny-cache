package handlers

import (
	"fmt"

	"github.com/nahK994/TinyCache/utils"
)

func HandleCommand(command string) error {
	fmt.Println(command)
	commandSegments, err := parseCommand(command)
	if err != nil {
		return err
	}

	commandType := commandSegments[0]
	if commandType == utils.GetCommand {
		fmt.Println("Get command found")
	} else if commandType == utils.SetCommand {
		fmt.Println("Set command found")
	}
	return nil
}
