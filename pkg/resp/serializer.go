package resp

import (
	"fmt"
	"strings"

	"github.com/nahK994/TinyCache/pkg/utils"
)

type CommandProcessor func([]string) string

var respCmd = utils.GetRESPCommands()
var commandProcessors = map[string]CommandProcessor{
	respCmd.SET:      processSET,
	respCmd.GET:      processGenericCommand,
	respCmd.EXISTS:   processGenericCommand,
	respCmd.INCR:     processGenericCommand,
	respCmd.DECR:     processGenericCommand,
	respCmd.DEL:      processGenericCommand,
	respCmd.LPUSH:    processGenericCommand,
	respCmd.LPOP:     processGenericCommand,
	respCmd.RPUSH:    processGenericCommand,
	respCmd.RPOP:     processGenericCommand,
	respCmd.LRANGE:   processGenericCommand,
	respCmd.FLUSHALL: processFlushAll,
	respCmd.PING:     processPing,
}

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

func getRESPformat(segments []string) string {
	serializedCmd := fmt.Sprintf("*%d\r\n", len(segments))

	serializedCmd += fmt.Sprintf("$%d\r\n%s\r\n", len(segments[0]), strings.ToUpper(segments[0]))
	for i := 1; i < len(segments); i++ {
		serializedCmd += fmt.Sprintf("$%d\r\n%s\r\n", len(segments[i]), segments[i])
	}
	return serializedCmd
}

func processSET(words []string) string {
	words = []string{words[0], words[1], strings.Join(words[2:], " ")}
	return getRESPformat(words)
}

func processGenericCommand(words []string) string {
	return getRESPformat(words)
}

func processFlushAll([]string) string {
	return "*1\r\n$8\r\nFLUSHALL\r\n"
}

func processPing([]string) string {
	return "*1\r\n$4\r\nPING\r\n"
}

func Serialize(rawCmd string) string {
	words := utils.GetCmdSegments(rawCmd)
	commandName := getCommandName(rawCmd)

	if processor, exists := commandProcessors[commandName]; exists {
		return processor(words)
	}
	return ""
}
