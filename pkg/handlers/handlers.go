package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nahK994/TinyCache/pkg/cache"
	"github.com/nahK994/TinyCache/pkg/errors"
	"github.com/nahK994/TinyCache/pkg/resp"
	"github.com/nahK994/TinyCache/pkg/utils"
)

var c *cache.Cache = cache.InitCache()
var errType = errors.GetErrorTypes()

func handleGET(key string) string {
	replytype := utils.GetReplyTypes()
	if !c.IsKeyExist(key) {
		return "$-1\r\n"
	}

	if val_int, ok_int := c.ReadCache(key).(int); ok_int {
		str := strconv.Itoa(val_int)
		return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(str), str)
	} else {
		val_str, _ := c.ReadCache(key).(string)
		return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(val_str), val_str)
	}
}

func handleSET(arguments []string) string {
	key := arguments[0]
	value := arguments[1]
	c.WriteCache(key, value)
	return "+OK\r\n"
}

func handleKeyExist(key string) string {
	if c.IsKeyExist(key) {
		return ":1\r\n"
	} else {
		return ":0\r\n"
	}
}

func handleINCR(key string) (string, error) {
	if !c.IsKeyExist(key) {
		c.WriteCache(key, 1)
		return ":1\r\n", nil
	} else {
		_, ok := c.ReadCache(key).(int)
		if !ok {
			return "", errors.Err{Type: errType.TypeError}
		}
		return c.INCRCache(key), nil
	}
}

func handleDECR(key string) (string, error) {
	if !c.IsKeyExist(key) {
		c.WriteCache(key, -1)
		return ":-1\r\n", nil
	} else {
		_, ok := c.ReadCache(key).(int)
		if !ok {
			return "", errors.Err{Type: errType.TypeError}
		}
		return c.DECRCache(key), nil
	}
}

func handleDEL(key string) string {
	if c.IsKeyExist(key) {
		c.DeleteCache(key)
		return ":1\r\n"
	} else {
		return ":0\r\n"
	}
}

func HandleCommand(serializedRawCmd string) (string, error) {
	cmdSegments := resp.Deserializer(serializedRawCmd)
	respCmd := utils.GetRESPCommands()

	cmd := cmdSegments[0]
	args := cmdSegments[1:]

	switch strings.ToUpper(cmd) {
	case respCmd.GET:
		return handleGET(args[0]), nil
	case respCmd.SET:
		return handleSET(args), nil
	case respCmd.EXISTS:
		return handleKeyExist(args[0]), nil
	case respCmd.DEL:
		return handleDEL(args[0]), nil
	case respCmd.PING:
		return "+PONG\r\n", nil
	case respCmd.INCR:
		return handleINCR(args[0])
	case respCmd.DECR:
		return handleDECR(args[0])
	default:
		return "", nil
	}
}
