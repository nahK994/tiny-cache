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

func handleGET(arguments []string) (string, error) {
	replytype := utils.GetReplyTypes()
	if len(arguments) > 1 {
		return "", errors.Err{Msg: "-ERR wrong number of arguments for 'GET' command\r\n", File: "handlers/handlers.go", Line: 19}
	}
	if !c.IsKeyExist(arguments[0]) {
		return "$-1\r\n", nil
	}

	val_str, ok_str := c.ReadCache(arguments[0]).(string)
	val_int, ok_int := c.ReadCache(arguments[0]).(int)
	if ok_int {
		str := strconv.Itoa(val_int)
		return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(str), str), nil
	} else if ok_str {
		return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(val_str), val_str), nil
	} else {
		return "", errors.Err{Msg: "-Unknown datatype\r\n", File: "handlers/handlers.go", Line: 32}
	}
}

func handleSET(arguments []string) (string, error) {
	key := arguments[0]
	value := arguments[1]
	if err := c.WriteCache(key, value); err != nil {
		return "", err
	}
	return "+OK\r\n", nil
}

func handleKeyExist(arguments []string) (string, error) {
	if len(arguments) != 1 {
		return "", errors.Err{Msg: "-ERR unknown command 'INVALID_COMMAND'\r\n", File: "handlers/handlers.go", Line: 42}
	}

	if c.IsKeyExist(arguments[0]) {
		return ":1\r\n", nil
	} else {
		return ":0\r\n", nil
	}
}

func handleINCR(arguments []string) (string, error) {
	if len(arguments) != 1 {
		return "", errors.Err{Msg: "-ERR unknown command 'INVALID_COMMAND'\r\n", File: "handlers/handlers.go", Line: 54}
	}

	if !c.IsKeyExist(arguments[0]) {
		c.WriteCache(arguments[0], 1)
		return ":1\r\n", nil
	} else {
		val, ok := c.ReadCache(arguments[0]).(int)
		if !ok {
			return "", errors.Err{Msg: "-ERR value aren't available for INCR\r\n", File: "handlers/handlers.go", Line: 63}
		}

		val++
		c.WriteCache(arguments[0], val)
		return fmt.Sprintf(":%d\r\n", val), nil
	}
}

func handleDECR(arguments []string) (string, error) {
	if len(arguments) != 1 {
		return "", errors.Err{Msg: "-ERR unknown command 'INVALID_COMMAND'\r\n", File: "handlers/handlers.go", Line: 75}
	}

	if !c.IsKeyExist(arguments[0]) {
		c.WriteCache(arguments[0], -1)
		return ":-1\r\n", nil
	} else {
		val, ok := c.ReadCache(arguments[0]).(int)
		if !ok {
			return "", errors.Err{Msg: "-ERR value aren't available for INCR\r\n", File: "handlers/handlers.go", Line: 84}
		}

		val--
		c.WriteCache(arguments[0], val)
		return fmt.Sprintf(":%d\r\n", val), nil
	}
}

func HandleCommand(serializedRawCmd string) (string, error) {
	cmdSegments, err := resp.Deserializer(serializedRawCmd)
	respCmd := utils.GetRESPCommands()
	if err != nil {
		return "", err
	}
	cmd := cmdSegments[0]
	args := cmdSegments[1:]

	switch strings.ToUpper(cmd) {
	case respCmd.GET:
		return handleGET(args)
	case respCmd.SET:
		return handleSET(args)
	case respCmd.EXISTS:
		return handleKeyExist(args)
	case respCmd.INCR:
		return handleINCR(args)
	case respCmd.DECR:
		return handleDECR(args)
	case respCmd.PING:
		return "+PONG\r\n", nil
	default:
		return fmt.Sprintln("Please use these commands:", strings.Join([]string{
			respCmd.SET, respCmd.GET, respCmd.EXISTS, respCmd.PING,
		}, ", ")), nil
	}
}
