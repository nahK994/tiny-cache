package handlers

import (
	"strconv"
	"strings"
	"time"

	"github.com/nahK994/TinyCache/pkg/cache"
	"github.com/nahK994/TinyCache/pkg/config"
	"github.com/nahK994/TinyCache/pkg/errors"
	"github.com/nahK994/TinyCache/pkg/resp"
	"github.com/nahK994/TinyCache/pkg/shared"
	"github.com/nahK994/TinyCache/pkg/validators"
)

func isKeyExists(key string) bool {
	_ = validators.ValidateExpiry(key)
	return c.EXISTS(key)
}

var c *cache.Cache = config.App.Cache

func handleGET(key string) (string, error) {
	if err := validators.AssertKeyExists(key); err != nil {
		return "", err
	}
	item := c.GET(key)
	return resp.SerializeCacheItem(item.Value), nil
}

func handleSET(key string, args []string) (string, error) {
	if len(args) > 1 {
		c.SET(key, args[0])
		ttl, _ := strconv.Atoi(args[1])
		c.EXPIRE(key, ttl)
	} else {
		c.SET(key, args[0])
	}
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
	return resp.SerializeBool(isKeyExists(key))
}

func handleIncDec(key, operation string) (string, error) {
	if !isKeyExists(key) {
		c.SET(key, 0)
	} else if c.GET(key).Value.DataType != cache.Int {
		return "", errors.Err{Type: errors.TypeError}
	}

	result := cache.CacheData{
		DataType: cache.Int,
	}
	switch operation {
	case resp.INCR:
		result.IntData = shared.IntToPtr(c.INCR(key))
	case resp.DECR:
		result.IntData = shared.IntToPtr(c.DECR(key))
	}

	return resp.SerializeCacheItem(result), nil
}

func handleINCR(key string) (string, error) {
	return handleIncDec(key, resp.INCR)
}

func handleDECR(key string) (string, error) {
	return handleIncDec(key, resp.DECR)
}

func handleDEL(key string) string {
	keyExists := isKeyExists(key)
	if keyExists {
		c.DEL(key)
	}
	return resp.SerializeBool(keyExists)
}

func handleLpushRpush(key string, args []string, operation string) (string, error) {
	if isKeyExists(key) && c.GET(key).Value.DataType != cache.Array {
		return "", errors.Err{Type: errors.TypeError}
	}

	switch operation {
	case resp.LPUSH:
		c.LPUSH(key, args)
	case resp.RPUSH:
		c.RPUSH(key, args)
	}

	item := cache.CacheData{
		DataType: cache.Int,
	}
	item.IntData = shared.IntToPtr(len(c.LRANGE(key, 0, -1)))
	return resp.SerializeCacheItem(item), nil
}

func handleLPUSH(key string, args []string) (string, error) {
	return handleLpushRpush(key, args, resp.LPUSH)
}

func handleRPUSH(key string, args []string) (string, error) {
	return handleLpushRpush(key, args, resp.RPUSH)
}

func handleLRANGE(key string, startIdx, endIdx int) (string, error) {
	if err := validators.AssertKeyExists(key); err != nil {
		return "", err
	}
	if c.GET(key).Value.DataType != cache.Array {
		return "", errors.Err{Type: errors.TypeError}
	}
	vals := c.LRANGE(key, startIdx, endIdx)
	item := cache.CacheData{
		DataType: cache.Array,
		StrList:  vals,
	}
	return resp.SerializeCacheItem(item), nil
}

func handleListPop(key, popType string) (string, error) {
	if err := validators.AssertKeyExists(key); err != nil {
		return "", err
	}

	item := c.GET(key).Value
	if item.DataType != cache.Array {
		return "", errors.Err{Type: errors.TypeError}
	}
	if len(item.StrList) <= 0 {
		return "", errors.Err{Type: errors.EmptyList}
	}

	cacheItem := cache.CacheData{
		DataType: cache.String,
	}
	switch popType {
	case resp.LPOP:
		cacheItem.StrData = &item.StrList[0]
		c.LPOP(key)
	case resp.RPOP:
		cacheItem.StrData = &item.StrList[len(item.StrList)-1]
		c.RPOP(key)
	}

	return resp.SerializeCacheItem(cacheItem), nil
}

func handleLPOP(key string) (string, error) {
	return handleListPop(key, resp.LPOP)
}

func handleRPOP(key string) (string, error) {
	return handleListPop(key, resp.RPOP)
}

func handleEXPIRE(key string, ttl int) (string, error) {
	if err := validators.AssertKeyExists(key); err != nil {
		return "", err
	}

	c.EXPIRE(key, ttl)
	return "+OK\r\n", nil
}

func handleTTL(key string) (string, error) {
	if err := validators.AssertKeyExists(key); err != nil {
		return "", err
	}
	cacheItem := cache.CacheData{
		DataType: cache.Int,
		IntData:  new(int),
	}

	item := c.GET(key)
	if item.ExpiryTime == nil {
		cacheItem.IntData = shared.IntToPtr(0)
	} else {
		if remainingTTL := int(time.Until(*item.ExpiryTime).Seconds()); remainingTTL > 0 {
			cacheItem.IntData = &remainingTTL
		} else {
			cacheItem.IntData = shared.IntToPtr(-1)
		}
	}
	return resp.SerializeCacheItem(cacheItem), nil
}

func handlePERSIST(key string) (string, error) {
	if err := validators.AssertKeyExists(key); err != nil {
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
		return handleSET(args[0], args[1:])
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
		ttl, _ := strconv.Atoi(args[1])
		return handleEXPIRE(args[0], ttl)
	case resp.LRANGE:
		startIdx, _ := strconv.Atoi(args[1])
		endIdx, _ := strconv.Atoi(args[2])
		return handleLRANGE(args[0], startIdx, endIdx)
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
		return "", errors.Err{Type: errors.UnknownCommand}
	}
}
