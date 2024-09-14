package resp

import (
	"fmt"
	"strings"

	"github.com/nahK994/TinyCache/pkg/utils"
)

func getCommandName(cmd string) string {
	seg := ""
	for _, ch := range cmd {
		if ch == ' ' {
			break
		}
		seg += string(ch)
	}
	return strings.ToUpper(seg)
}

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

	serializedCmd += fmt.Sprintf("$%d\r\n%s\r\n", len(segments[0]), strings.ToUpper(segments[0]))
	for i := 1; i < len(segments); i++ {
		serializedCmd += fmt.Sprintf("$%d\r\n%s\r\n", len(segments[i]), segments[i])
	}
	return serializedCmd
}

func processSET(cmd string) string {
	words := getCmdSegments(cmd)
	words = []string{words[0], words[1], strings.Join(words[2:], " ")}
	return getRESPformat(words)
}

func processKeyBasedCommand(cmd string) string {
	words := getCmdSegments(cmd)
	return getRESPformat(words)
}

func Serialize(rawCmd string) string {
	respCmd := utils.GetRESPCommands()

	switch getCommandName(rawCmd) {
	case respCmd.SET:
		return processSET(rawCmd)
	case respCmd.GET:
		return processKeyBasedCommand(rawCmd)
	case respCmd.EXISTS:
		return processKeyBasedCommand(rawCmd)
	case respCmd.INCR:
		return processKeyBasedCommand(rawCmd)
	case respCmd.DECR:
		return processKeyBasedCommand(rawCmd)
	case respCmd.DEL:
		return processKeyBasedCommand(rawCmd)
	case respCmd.PING:
		return "*1\r\n$4\r\nPING\r\n"
	}
	return ""
}
