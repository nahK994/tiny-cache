package resp

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/nahK994/TinyCache/pkg/cache"
	"github.com/nahK994/TinyCache/pkg/shared"
	"github.com/nahK994/TinyCache/pkg/utils"
)

func processArray(segments []string) string {
	response := fmt.Sprintf("*%d\r\n", len(segments))
	for _, v := range segments {
		response += fmt.Sprintf("$%d\r\n%s\r\n", len(v), v)
	}
	return response
}

func Serialize(rawCmd string) string {
	words := shared.SplitCmd(rawCmd)
	commandName := strings.ToUpper(words[0])

	if commandName == PING {
		return "*1\r\n$4\r\nPING\r\n"
	} else if commandName == FLUSHALL {
		return "*1\r\n$8\r\nFLUSHALL\r\n"
	} else {
		return processArray(words)
	}
}

func SerializeBool(arg bool) string {
	if arg {
		return ":1\r\n"
	} else {
		return ":0\r\n"
	}
}

func SerializeCacheItem(item cache.DataItem) string {
	switch item.DataType {
	case utils.Int:
		val, _ := strconv.Atoi(string(item.Value))
		return fmt.Sprintf(":%d\r\n", val)
	case utils.String:
		val := string(item.Value)
		return fmt.Sprintf("$%d\r\n%s\r\n", len(val), val)
	case utils.Array:
		var vals []string
		json.Unmarshal(item.Value, &vals)
		return processArray(vals)
	}
	return "$-1\r\n"
}
