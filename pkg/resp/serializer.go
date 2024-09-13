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

func processSET(cmd string) (string, error) {
	words := getCmdSegments(cmd)
	if len(words) < 3 {
		return "", errors.Err{Msg: "invalid SET command argument", File: "resp/serializer.go", Line: 34}
	}

	words = []string{words[0], words[1], strings.Join(words[2:], " ")}
	return getRESPformat(words), nil
}

func processGET(cmd string) (string, error) {
	words := getCmdSegments(cmd)
	if len(words) != 2 {
		return "", errors.Err{Msg: "invalid GET command argument", File: "resp/serializer.go", Line: 47}
	}

	return getRESPformat(words), nil
}

func processEXISTS(cmd string) (string, error) {
	words := getCmdSegments(cmd)
	if len(words) != 2 {
		return "", errors.Err{Msg: "invalid EXISTS command argument", File: "resp/serializer.go", Line: 56}
	}

	return getRESPformat(words), nil
}

func processINCR(cmd string) (string, error) {
	words := getCmdSegments(cmd)
	if len(words) != 2 {
		return "", errors.Err{Msg: "invalid INCR command argument", File: "resp/serializer.go", Line: 73}
	}

	return getRESPformat(words), nil
}

func processDECR(cmd string) (string, error) {
	words := getCmdSegments(cmd)
	if len(words) != 2 {
		return "", errors.Err{Msg: "invalid DECR command argument", File: "resp/serializer.go", Line: 82}
	}

	return getRESPformat(words), nil
}

func Serialize(rawCmd string) (string, error) {
	respCmd := utils.GetRESPCommands()

	switch strings.ToUpper(getCommandName(rawCmd)) {
	case respCmd.SET:
		return processSET(rawCmd)
	case respCmd.GET:
		return processGET(rawCmd)
	case respCmd.EXISTS:
		return processEXISTS(rawCmd)
	case respCmd.INCR:
		return processINCR(rawCmd)
	case respCmd.DECR:
		return processDECR(rawCmd)
	case respCmd.PING:
		return "*1\r\n$4\r\nPING\r\n", nil
	default:
		return "", errors.Err{Msg: utils.GetClientMessage(), File: "resp/serializer.go", Line: 82}
	}
}
