package handlers

import (
	"fmt"
)

func HandleCommand(msg []byte) error {
	fmt.Println(string(msg))
	return nil
}
