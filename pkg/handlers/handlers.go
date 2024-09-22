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
var replytype = utils.GetReplyTypes()
var errType = errors.GetErrorTypes()

func handleGET(key string) string {
	if !c.EXISTS(key) {
		return fmt.Sprintf("%c-1\r\n", replytype.Int)
	}

	if val_int, ok_int := c.GET(key).(int); ok_int {
		str := strconv.Itoa(val_int)
		return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(str), str)
	} else {
		val_str, _ := c.GET(key).(string)
		return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(val_str), val_str)
	}
}

func handleSET(arguments []string) string {
	key := arguments[0]
	value := arguments[1]
	c.SET(key, value)
	return "+OK\r\n"
}

func handleKeyExist(key string) string {
	if c.EXISTS(key) {
		return fmt.Sprintf("%c1\r\n", replytype.Int)
	} else {
		return fmt.Sprintf("%c0\r\n", replytype.Int)
	}
}

func handleINCR(key string) (string, error) {
	if !c.EXISTS(key) {
		c.SET(key, 1)
		return fmt.Sprintf("%c1\r\n", replytype.Int), nil
	} else {
		_, ok := c.GET(key).(int)
		if !ok {
			return "", errors.Err{Type: errType.TypeError}
		}
		return fmt.Sprintf("%c%d\r\n", replytype.Int, c.INCR(key)), nil
	}
}

func handleDECR(key string) (string, error) {
	if !c.EXISTS(key) {
		c.SET(key, -1)
		return fmt.Sprintf("%c-1\r\n", replytype.Int), nil
	} else {
		_, ok := c.GET(key).(int)
		if !ok {
			return "", errors.Err{Type: errType.TypeError}
		}
		return fmt.Sprintf("%c%d\r\n", replytype.Int, c.DECR(key)), nil
	}
}

func handleDEL(key string) string {
	if c.EXISTS(key) {
		c.DEL(key)
		return fmt.Sprintf("%c1\r\n", replytype.Int)
	} else {
		return fmt.Sprintf("%c0\r\n", replytype.Int)
	}
}

func handleLPUSH(key string, args []string) string {
	c.LPUSH(key, args)
	vals := c.LRANGE(key, 1, -1)
	return fmt.Sprintf("%c%d\r\n", replytype.Int, len(vals))
}

func handleLRANGE(key string, startIdx, endIdx int) string {
	vals := c.LRANGE(key, startIdx, endIdx)
	var response string
	response += fmt.Sprintf("%c%d\r\n", replytype.Array, len(vals))
	for i := 0; i < len(vals); i++ {
		response += fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(vals[i]), vals[i])
	}
	return response
}

func handleLPOP(key string) string {
	val := c.LRANGE(key, 1, 1)
	if len(val) > 0 {
		c.LPOP(key)
		return fmt.Sprintf("%c1\r\n%c%d\r\n%s\r\n", replytype.Array, replytype.Bulk, len(val), val[0])
	} else {
		return fmt.Sprintf("%c0\r\n", replytype.Array)
	}
}

func HandleCommand(serializedRawCmd string) string {
	cmdSegments, _ := resp.Deserializer(serializedRawCmd).([]string)
	respCmd := utils.GetRESPCommands()

	cmd := cmdSegments[0]
	args := cmdSegments[1:]

	switch strings.ToUpper(cmd) {
	case respCmd.GET:
		return handleGET(args[0])
	case respCmd.SET:
		return handleSET(args)
	case respCmd.EXISTS:
		return handleKeyExist(args[0])
	case respCmd.DEL:
		return handleDEL(args[0])
	case respCmd.PING:
		return "+PONG\r\n"
	case respCmd.LPUSH:
		return handleLPUSH(args[0], args[1:])
	case respCmd.LRANGE:
		strIdx, _ := strconv.Atoi(args[1])
		endIdx, _ := strconv.Atoi(args[2])
		return handleLRANGE(args[0], strIdx, endIdx)
	case respCmd.LPOP:
		return handleLPOP(args[0])
	case respCmd.INCR:
		val, err := handleINCR(args[0])
		if err != nil {
			return err.Error()
		}
		return val
	case respCmd.DECR:
		val, err := handleDECR(args[0])
		if err != nil {
			return err.Error()
		}
		return val
	default:
		return ""
	}
}
