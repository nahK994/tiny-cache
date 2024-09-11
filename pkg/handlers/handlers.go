package handlers

import (
	"fmt"
	"strings"

	"github.com/nahK994/TinyCache/pkg/resp"
	"github.com/nahK994/TinyCache/pkg/utils"
)

func handleGET(cmdSegments []string) (string, error) {
	return "Processing GET command ......", nil
}

func handleSET(cmdSegments []string) (string, error) {
	return "Processing SET command ......", nil
}

func HandleCommand(serializedCmd string) (string, error) {
	cmdSegments, err := resp.Deserializer(serializedCmd)
	respCmd := utils.GetRESPCommands()
	if err != nil {
		return "", err
	}

	switch cmdSegments[0] {
	case respCmd.GET:
		return handleGET(cmdSegments)
	case respCmd.SET:
		return handleSET(cmdSegments)
	default:
		return fmt.Sprintln("Please use these commands:", strings.Join([]string{
			respCmd.SET, respCmd.GET,
		}, ", ")), nil
	}

	// return "+OK\r\n", nil
}
