package resp

import (
	"fmt"
	"strings"

	"github.com/nahK994/TinyCache/pkg/cache"
	"github.com/nahK994/TinyCache/pkg/shared"
)

func processArray(segments []string) string {
	response := fmt.Sprintf("*%d\r\n", len(segments))
	for _, v := range segments {
		response += fmt.Sprintf("$%d\r\n%s\r\n", len(v), v)
	}
	return response
}

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

func SerializeBool(arg bool) string {
	if arg {
		return ":1\r\n"
	} else {
		return ":0\r\n"
	}
}

func SerializeCacheItem(item cache.CacheData) string {
	switch item.DataType {
	case cache.Int:
		if item.IntData != nil {
			return fmt.Sprintf(":%d\r\n", *item.IntData)
		}
	case cache.String:
		if item.StrData != nil {
			return fmt.Sprintf("$%d\r\n%s\r\n", len(*item.StrData), *item.StrData)
		}
	case cache.Array:
		if item.StrList != nil {
			response := fmt.Sprintf("*%d\r\n", len(item.StrList))
			for _, v := range item.StrList {
				response += fmt.Sprintf("$%d\r\n%s\r\n", len(v), v)
			}
			return processArray(item.StrList)
		}
	}
	return "$-1\r\n"
}
