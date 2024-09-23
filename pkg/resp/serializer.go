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
	respCmd.LRANGE:   processGenericCommand,
	respCmd.LPOP:     processGenericCommand,
	respCmd.FLUSHALL: processFlushAll,
	respCmd.PING:     processPing,
}

func getCommandName(cmd string) string {
	seg := strings.Split(cmd, " ")[0]
	return strings.ToUpper(seg)
}

func getCmdSegments(cmd string) []string {
	return strings.Fields(cmd)
}

func getRESPformat(segments []string) string {
	serializedCmd := fmt.Sprintf("*%d\r\n", len(segments))
	for _, seg := range segments {
		serializedCmd += fmt.Sprintf("$%d\r\n%s\r\n", len(seg), seg)
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
	words := getCmdSegments(rawCmd)
	commandName := getCommandName(rawCmd)

	if processor, exists := commandProcessors[commandName]; exists {
		return processor(words)
	}
	return ""
}
