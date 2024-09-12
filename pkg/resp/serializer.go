package resp

import (
	"fmt"
	"strings"

	"github.com/nahK994/TinyCache/pkg/errors"
	"github.com/nahK994/TinyCache/pkg/utils"
)

func getCmdSegments(cmd string) []string {
	var words []string
	temp := strings.Split(cmd, " ")
	for _, ch := range temp {
		if len(ch) == 0 {
			continue
		}
		words = append(words, ch)
	}
	return words
}

func getRESPformat(segments []string) string {
	serializedCmd := fmt.Sprintf("*%d\r\n", len(segments))
	for i := 0; i < len(segments); i++ {
		serializedCmd += fmt.Sprintf("$%d\r\n%s\r\n", len(segments[i]), segments[i])
	}
	return serializedCmd
}

func getCommandName(cmd string) string {
	seg := ""
	for _, ch := range cmd {
		if ch == ' ' {
			break
		}
		seg += string(ch)
	}
	return seg
}

func processSETcommand(cmd string) (string, error) {
	words := getCmdSegments(cmd)
	if len(words) < 3 {
		return "", errors.Err{Msg: "invalid SET command argument", File: "resp/serializer.go", Line: 34}
	}

	words = []string{words[0], words[1], strings.Join(words[2:], " ")}
	return getRESPformat(words), nil
}

func processGETcommand(cmd string) (string, error) {
	words := getCmdSegments(cmd)
	if len(words) != 2 {
		return "", errors.Err{Msg: "invalid GET command argument", File: "resp/serializer.go", Line: 47}
	}

	return getRESPformat(words), nil
}

func processEXISTScommand(cmd string) (string, error) {
	words := getCmdSegments(cmd)
	if len(words) != 2 {
		return "", errors.Err{Msg: "invalid EXISTS command argument", File: "resp/serializer.go", Line: 56}
	}

	return getRESPformat(words), nil
}

func Serialize(cmd string) (string, error) {
	respCmd := utils.GetRESPCommands()

	switch strings.ToUpper(getCommandName(cmd)) {
	case respCmd.SET:
		return processSETcommand(cmd)
	case respCmd.GET:
		return processGETcommand(cmd)
	case respCmd.EXISTS:
		return processEXISTScommand(cmd)
	default:
		return "", errors.Err{Msg: fmt.Sprintln("Please use these commands:", strings.Join([]string{
			respCmd.SET, respCmd.GET, respCmd.EXISTS,
		}, ", ")), File: "resp/serializer.go", Line: 82}
	}
}
