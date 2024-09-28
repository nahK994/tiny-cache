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
)

var c *cache.Cache = config.App.Cache

// var replytype = utils.GetReplyTypes()

func handleGET(key string) (string, error) {
	if err := AssertKeyExists(key); err != nil {
		return "", err
	}
	item := c.GET(key)
	val := item.Value
	switch val.DataType {
	case cache.Int:
		str := strconv.Itoa(*val.IntData)
		return fmt.Sprintf("$%d\r\n%s\r\n", len(str), str), nil
	case cache.String:
		return fmt.Sprintf("$%d\r\n%s\r\n", len(*val.StrData), *val.StrData), nil
	default:
		return "", errors.Err{Type: errors.TypeError}
	}
}

func handleSET(key, value string) (string, error) {
	c.SET(key, value)
	return "+OK\r\n", nil
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
	if IsKeyExists(key) {
		return ":1\r\n"
	} else {
		return ":0\r\n"
	}
}

func handleIncDec(key string, operation string) (string, error) {
	if !IsKeyExists(key) {
		c.SET(key, 0)
	} else if c.GET(key).Value.DataType != cache.Int {
		return "", errors.Err{Type: errors.TypeError}
	}

	var result int
	switch operation {
	case resp.INCR:
		result = c.INCR(key)
	case resp.DECR:
		result = c.DECR(key)
	default:
		return "", errors.Err{Type: errors.UnknownCommand}
	}

	return fmt.Sprintf(":%d\r\n", result), nil
}

func handleINCR(key string) (string, error) {
	return handleIncDec(key, resp.INCR)
}

func handleDECR(key string) (string, error) {
	return handleIncDec(key, resp.DECR)
}

func handleDEL(key string) string {
	if IsKeyExists(key) {
		c.DEL(key)
		return ":1\r\n"
	} else {
		return ":0\r\n"
	}
}

func handleLPUSH(key string, args []string) (string, error) {
	if IsKeyExists(key) {
		if c.GET(key).Value.DataType != cache.Array {
			return "", errors.Err{Type: errors.TypeError}
		}
	}

	c.LPUSH(key, args)
	vals := c.LRANGE(key, 0, -1)
	return fmt.Sprintf(":%d\r\n", len(vals)), nil
}

func handleRPUSH(key string, args []string) (string, error) {
	if IsKeyExists(key) {
		if c.GET(key).Value.DataType != cache.Array {
			return "", errors.Err{Type: errors.TypeError}
		}
	}

	c.RPUSH(key, args)
	vals := c.LRANGE(key, 0, -1)
	return fmt.Sprintf(":%d\r\n", len(vals)), nil
}

func handleLRANGE(key string, startIdx, endIdx int) (string, error) {
	if err := AssertKeyExists(key); err != nil {
		return "", err
	}
	if c.GET(key).Value.DataType != cache.Array {
		return "", errors.Err{Type: errors.TypeError}
	}
	vals := c.LRANGE(key, startIdx, endIdx)
	response := fmt.Sprintf("*%d\r\n", len(vals))
	for i := 0; i < len(vals); i++ {
		response += fmt.Sprintf("$%d\r\n%s\r\n", len(vals[i]), vals[i])
	}
	return response, nil
}

func handleListPop(key string, popType string) (string, error) {
	if err := AssertKeyExists(key); err != nil {
		return "", err
	}
	if c.GET(key).Value.DataType != cache.Array {
		return "", errors.Err{Type: errors.TypeError}
	}

	var val []string
	switch popType {
	case resp.LPOP:
		val = c.LRANGE(key, 0, 0)
		if len(val) > 0 {
			data := val[0]
			c.LPOP(key)
			return fmt.Sprintf("$%d\r\n%s\r\n", len(data), data), nil
		}
	case resp.RPOP:
		val = c.LRANGE(key, 0, -1)
		if len(val) > 0 {
			data := val[len(val)-1]
			c.RPOP(key)
			return fmt.Sprintf("$%d\r\n%s\r\n", len(data), data), nil
		}
	}
	return "$0\r\n", nil
}

func handleLPOP(key string) (string, error) {
	if err := CheckEmptyList(key); err != nil {
		return "", err
	}
	if c.GET(key).Value.DataType != cache.Array {
		return "", errors.Err{Type: errors.TypeError}
	}
	return handleListPop(key, resp.LPOP)
}

func handleRPOP(key string) (string, error) {
	if err := CheckEmptyList(key); err != nil {
		return "", err
	}
	if c.GET(key).Value.DataType != cache.Array {
		return "", errors.Err{Type: errors.TypeError}
	}
	return handleListPop(key, resp.RPOP)
}

func handleEXPIRE(key string, ttl int) (string, error) {
	if err := AssertKeyExists(key); err != nil {
		return "", err
	}

	if ttl < 0 {
		return "", errors.Err{Type: errors.InvalidCommandFormat}
	}
	c.EXPIRE(key, ttl)
	return "+OK\r\n", nil
}

func handleTTL(key string) (string, error) {
	if err := AssertKeyExists(key); err != nil {
		return "", err
	}

	item := c.GET(key)
	if item.ExpiryTime == nil {
		return ":0\r\n", nil
	}

	remainingTTL := int(time.Until(*item.ExpiryTime).Seconds())
	if remainingTTL < 0 {
		return ":-1\r\n", nil
	}

	return fmt.Sprintf(":%d\r\n", remainingTTL), nil
}

func handlePERSIST(key string) (string, error) {
	if err := AssertKeyExists(key); err != nil {
		return "", err
	}

	c.EXPIRE(key, 0)
	return "+OK\r\n", nil
}

func HandleCommand(serializedRawCmd string) (string, error) {
	cmdSegments, _ := resp.Deserializer(serializedRawCmd).([]string)

	cmd := cmdSegments[0]
	args := cmdSegments[1:]

	switch strings.ToUpper(cmd) {
	case resp.GET:
		return handleGET(args[0])
	case resp.SET:
		return handleSET(args[0], args[1])
	case resp.EXISTS:
		return handleEXISTS(args[0]), nil
	case resp.DEL:
		return handleDEL(args[0]), nil
	case resp.PING:
		return "+PONG\r\n", nil
	case resp.LPUSH:
		return handleLPUSH(args[0], args[1:])
	case resp.RPUSH:
		return handleRPUSH(args[0], args[1:])
	case resp.EXPIRE:
		key := args[0]
		ttl, _ := strconv.Atoi(args[1])
		return handleEXPIRE(key, ttl)
	case resp.LRANGE:
		strIdx, _ := strconv.Atoi(args[1])
		endIdx, _ := strconv.Atoi(args[2])
		return handleLRANGE(args[0], strIdx, endIdx)
	case resp.LPOP:
		return handleLPOP(args[0])
	case resp.RPOP:
		return handleRPOP(args[0])
	case resp.INCR:
		return handleINCR(args[0])
	case resp.DECR:
		return handleDECR(args[0])
	case resp.FLUSHALL:
		return handleFLUSHALL(), nil
	case resp.TTL:
		return handleTTL(args[0])
	case resp.PERSIST:
		return handlePERSIST(args[0])
	default:
		return "", nil
	}
}
