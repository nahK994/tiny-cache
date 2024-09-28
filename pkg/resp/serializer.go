package resp

import (
	"strings"

	"github.com/nahK994/TinyCache/pkg/shared"
)

type CommandProcessor func([]string) string

var commandProcessors = map[string]CommandProcessor{
	SET:      processSET,
	GET:      processGenericCommand,
	EXISTS:   processGenericCommand,
	INCR:     processGenericCommand,
	DECR:     processGenericCommand,
	DEL:      processGenericCommand,
	LPUSH:    processGenericCommand,
	LPOP:     processGenericCommand,
	RPUSH:    processGenericCommand,
	RPOP:     processGenericCommand,
	LRANGE:   processGenericCommand,
	EXPIRE:   processGenericCommand,
	TTL:      processGenericCommand,
	PERSIST:  processGenericCommand,
	FLUSHALL: processFlushAll,
	PING:     processPing,
}

func processSET(words []string) string {
	words = []string{words[0], words[1], strings.Join(words[2:], " ")}
	return processArray(words)
}

func processGenericCommand(words []string) string {
	return processArray(words)
}

func processFlushAll([]string) string {
	return "*1\r\n$8\r\nFLUSHALL\r\n"
}

func processPing([]string) string {
	return "*1\r\n$4\r\nPING\r\n"
}

func Serialize(rawCmd string) string {
	words := shared.SplitCmd(rawCmd)
	commandName := strings.ToUpper(words[0])

	if processor, exists := commandProcessors[commandName]; exists {
		return processor(words)
	}
	return ""
}
