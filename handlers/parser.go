package handlers

import (
	"errors"
	"strings"

	"github.com/nahK994/TinyCache/utils"
)

func parseCommand(str string) ([]string, error) {
	commandLines := strings.Split(str, "\r\n")
	commandLines = commandLines[0 : len(commandLines)-1]

	var parsedCommandLines []string
	if len(commandLines) < 3 {
		return nil, errors.New("command cannot be parsed")
	}

	if strings.ToUpper(commandLines[2]) == utils.SetCommand {
		parsedCommandLines = append(parsedCommandLines, utils.SetCommand)
	} else if strings.ToUpper(commandLines[2]) == utils.GetCommand {
		parsedCommandLines = append(parsedCommandLines, utils.GetCommand)
	} else {
		return nil, errors.New("command cannot be parsed")
	}
	// fmt.Println(commandLines)
	// fmt.Println(parsedCommandLines)
	return parsedCommandLines, nil
}
