package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nahK994/TinyCache/pkg/cache"
	"github.com/nahK994/TinyCache/pkg/config"
	"github.com/nahK994/TinyCache/pkg/errors"
	"github.com/nahK994/TinyCache/pkg/resp"
	"github.com/nahK994/TinyCache/pkg/utils"
)

var c *cache.Cache = config.App.Cache
var replytype = utils.GetReplyTypes()
var errType = errors.GetErrorTypes()
var respCmd = utils.GetRESPCommands()

func checkExpirity(key string) (cache.Data, error) {
	item := c.GET(key)
	if item.ExpiryTime != nil && time.Now().After(*item.ExpiryTime) {
		c.DEL(key)
		return cache.Data{}, errors.Err{Type: errType.UndefinedKey}
	}
	return item, nil
}

func handleGET(key string) (string, error) {
	if !c.EXISTS(key) {
		return "", errors.Err{Type: errType.UndefinedKey}
	}

	item, err := checkExpirity(key)
	if err != nil {
		return "", err
	}

	switch val := item.Val.(type) {
	case int:
		str := strconv.Itoa(val)
		return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(str), str), nil
	case string:
		return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(val), val), nil
	default:
		return "", errors.Err{Type: errType.TypeError}
	}
}

func handleSET(key, value string) string {
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

func isKeyExists(key string) bool {
	checkExpirity(key)
	return c.EXISTS(key)
}

func handleEXISTS(key string) string {
	if isKeyExists(key) {
		return fmt.Sprintf("%c1\r\n", replytype.Int)
	} else {
		return fmt.Sprintf("%c0\r\n", replytype.Int)
	}
}

func handleIncDec(key string, operation string) (string, error) {
	checkExpirity(key)

	if !c.EXISTS(key) {
		c.SET(key, 0)
	} else {
		item := c.GET(key)
		_, ok := item.Val.(int)
		if !ok {
			return "", errors.Err{Type: errType.TypeError}
		}
	}

	var result int
	switch operation {
	case respCmd.INCR:
		result = c.INCR(key)
	case respCmd.DECR:
		result = c.DECR(key)
	default:
		return "", errors.Err{Type: errType.UnknownCommand}
	}

	return fmt.Sprintf("%c%d\r\n", replytype.Int, result), nil
}

func handleINCR(key string) (string, error) {
	return handleIncDec(key, respCmd.INCR)
}

func handleDECR(key string) (string, error) {
	return handleIncDec(key, respCmd.DECR)
}

func handleDEL(key string) string {
	checkExpirity(key)

	if c.EXISTS(key) {
		c.DEL(key)
		return fmt.Sprintf("%c1\r\n", replytype.Int)
	} else {
		return fmt.Sprintf("%c0\r\n", replytype.Int)
	}
}

func handleLPUSH(key string, args []string) (string, error) {
	typeErr := errors.Err{Type: errType.TypeError}
	err := validateListKey(key)
	if err.Error() == typeErr.Error() {
		return "", err
	}

	c.LPUSH(key, args)
	vals := c.LRANGE(key, 0, -1)
	return fmt.Sprintf("%c%d\r\n", replytype.Int, len(vals)), nil
}

func handleRPUSH(key string, args []string) (string, error) {
	typeErr := errors.Err{Type: errType.TypeError}
	err := validateListKey(key)
	if err.Error() == typeErr.Error() {
		return "", err
	}

	c.RPUSH(key, args)
	vals := c.LRANGE(key, 0, -1)
	return fmt.Sprintf("%c%d\r\n", replytype.Int, len(vals)), nil
}

func validateListKey(key string) error {
	if !isKeyExists(key) {
		return errors.Err{Type: errType.EmptyList}
	}

	if _, ok := c.GET(key).Val.([]string); !ok {
		return errors.Err{Type: errType.TypeError}
	}
	return nil
}

func handleLRANGE(key string, startIdx, endIdx int) (string, error) {
	if err := validateListKey(key); err != nil {
		return "", err
	}
	vals := c.LRANGE(key, startIdx, endIdx)
	response := fmt.Sprintf("%c%d\r\n", replytype.Array, len(vals))
	for i := 0; i < len(vals); i++ {
		response += fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(vals[i]), vals[i])
	}
	return response, nil
}

func handleListPop(key string, popType string) (string, error) {
	if err := validateListKey(key); err != nil {
		return "", err
	}

	var val []string
	switch popType {
	case respCmd.LPOP:
		val = c.LRANGE(key, 0, 0)
		if len(val) > 0 {
			data := val[0]
			c.LPOP(key)
			return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(data), data), nil
		}
	case respCmd.RPOP:
		val = c.LRANGE(key, 0, -1)
		if len(val) > 0 {
			data := val[len(val)-1]
			c.RPOP(key)
			return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(data), data), nil
		}
	}
	return fmt.Sprintf("%c0\r\n", replytype.Bulk), nil
}

func handleLPOP(key string) (string, error) {
	return handleListPop(key, respCmd.LPOP)
}

func handleRPOP(key string) (string, error) {
	return handleListPop(key, respCmd.RPOP)
}

func handleEXPIRE(key string, ttl int) (string, error) {
	if !isKeyExists(key) {
		return "", errors.Err{Type: errType.UndefinedKey}
	}

	if ttl < 0 {
		return "", errors.Err{Type: errType.InvalidCommandFormat}
	}
	c.EXPIRE(key, ttl)
	return "+OK\r\n", nil
}

func handleTTL(key string) (string, error) {
	item, notExistErr := checkExpirity(key)
	if notExistErr != nil {
		return "", notExistErr
	}

	if item.ExpiryTime == nil {
		return ":0\r\n", nil
	}

	remainingTTL := int(time.Until(*item.ExpiryTime).Seconds())
	if remainingTTL < 0 {
		return ":-1\r\n", nil
	}

	return fmt.Sprintf("%c%d\r\n", replytype.Int, remainingTTL), nil
}

func handlePERSIST(key string) (string, error) {
	if !isKeyExists(key) {
		return "", errors.Err{Type: errType.UndefinedKey}
	}

	c.EXPIRE(key, 0)
	return "+OK\r\n", nil
}

func HandleCommand(serializedRawCmd string) (string, error) {
	cmdSegments, _ := resp.Deserializer(serializedRawCmd).([]string)

	cmd := cmdSegments[0]
	args := cmdSegments[1:]

	switch strings.ToUpper(cmd) {
	case respCmd.GET:
		return handleGET(args[0])
	case respCmd.SET:
		return handleSET(args[0], args[1]), nil
	case respCmd.EXISTS:
		return handleEXISTS(args[0]), nil
	case respCmd.DEL:
		return handleDEL(args[0]), nil
	case respCmd.PING:
		return "+PONG\r\n", nil
	case respCmd.LPUSH:
		return handleLPUSH(args[0], args[1:])
	case respCmd.RPUSH:
		return handleRPUSH(args[0], args[1:])
	case respCmd.EXPIRE:
		key := args[0]
		ttl, _ := strconv.Atoi(args[1])
		return handleEXPIRE(key, ttl)
	case respCmd.LRANGE:
		strIdx, _ := strconv.Atoi(args[1])
		endIdx, _ := strconv.Atoi(args[2])
		return handleLRANGE(args[0], strIdx, endIdx)
	case respCmd.LPOP:
		return handleLPOP(args[0])
	case respCmd.RPOP:
		return handleRPOP(args[0])
	case respCmd.INCR:
		return handleINCR(args[0])
	case respCmd.DECR:
		return handleDECR(args[0])
	case respCmd.FLUSHALL:
		return handleFLUSHALL(), nil
	case respCmd.TTL:
		return handleTTL(args[0])
	case respCmd.PERSIST:
		return handlePERSIST(args[0])
	default:
		return "", nil
	}
}
