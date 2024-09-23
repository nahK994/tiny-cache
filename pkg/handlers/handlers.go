package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nahK994/TinyCache/pkg/cache"
	"github.com/nahK994/TinyCache/pkg/config"
	"github.com/nahK994/TinyCache/pkg/errors"
	"github.com/nahK994/TinyCache/pkg/resp"
	"github.com/nahK994/TinyCache/pkg/utils"
)

var c *cache.Cache = config.App.Cache
var replytype = utils.GetReplyTypes()
var errType = errors.GetErrorTypes()

func handleGET(key string) (string, error) {
	if !c.EXISTS(key) {
		return "", errors.Err{Type: errType.UndefinedKey}
	}

	if val_int, ok_int := c.GET(key).(int); ok_int {
		str := strconv.Itoa(val_int)
		return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(str), str), nil
	} else if val_str, ok_str := c.GET(key).(string); ok_str {
		return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(val_str), val_str), nil
	} else {
		return "", errors.Err{Type: errType.TypeError}
	}
}

func handleSET(arguments []string) string {
	key := arguments[0]
	value := arguments[1]
	c.SET(key, value)
	return "+OK\r\n"
}

func handleFLUSHALL() string {
	if config.App.IsAsyncFlush {
		config.App.FlushCh <- 1
	} else {
		c.FLUSHALL()
	}
	return "+OK\r\n"
}

func handleEXISTS(key string) string {
	if c.EXISTS(key) {
		return fmt.Sprintf("%c1\r\n", replytype.Int)
	} else {
		return fmt.Sprintf("%c0\r\n", replytype.Int)
	}
}

func handleINCR(key string) (string, error) {
	if !c.EXISTS(key) {
		c.SET(key, 0)
	} else {
		_, ok := c.GET(key).(int)
		if !ok {
			return "", errors.Err{Type: errType.TypeError}
		}
	}
	return fmt.Sprintf("%c%d\r\n", replytype.Int, c.INCR(key)), nil
}

func handleDECR(key string) (string, error) {
	if !c.EXISTS(key) {
		c.SET(key, 0)
	} else {
		_, ok := c.GET(key).(int)
		if !ok {
			return "", errors.Err{Type: errType.TypeError}
		}
	}
	return fmt.Sprintf("%c%d\r\n", replytype.Int, c.DECR(key)), nil
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
	vals := c.LRANGE(key, 0, -1)
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

func handleLPOP(key string) (string, error) {
	if !c.EXISTS(key) {
		return "", errors.Err{Type: errType.EmptyList}
	}
	_, err := handleGET(key)
	if err == nil {
		return "", errors.Err{Type: errType.TypeError}
	}

	val := c.LRANGE(key, 0, 0)
	if len(val) > 0 {
		c.LPOP(key)
		return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(val[0]), val[0]), nil
	} else {
		return fmt.Sprintf("%c0\r\n", replytype.Bulk), nil
	}
}

func HandleCommand(serializedRawCmd string) string {
	cmdSegments, _ := resp.Deserializer(serializedRawCmd).([]string)
	respCmd := utils.GetRESPCommands()

	cmd := cmdSegments[0]
	args := cmdSegments[1:]

	switch strings.ToUpper(cmd) {
	case respCmd.GET:
		val, err := handleGET(args[0])
		if err != nil {
			return err.Error()
		}
		return val
	case respCmd.SET:
		return handleSET(args)
	case respCmd.EXISTS:
		return handleEXISTS(args[0])
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
		val, err := handleLPOP(args[0])
		if err != nil {
			return err.Error()
		}
		return val
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
	case respCmd.FLUSHALL:
		return handleFLUSHALL()
	default:
		return ""
	}
}
