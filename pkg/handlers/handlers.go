package handlers

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/nahK994/TinyCache/pkg/cache"
	"github.com/nahK994/TinyCache/pkg/config"
	"github.com/nahK994/TinyCache/pkg/errors"
	"github.com/nahK994/TinyCache/pkg/resp"
	"github.com/nahK994/TinyCache/pkg/utils"
	"github.com/nahK994/TinyCache/pkg/validators"
)

var c *cache.Cache = config.App.Cache

func handleGET(key string) (string, error) {
	if !c.EXISTS(key) {
		return "", errors.Err{Type: errors.UndefinedKey}
	}

	if err := validators.ValidateExpiry(key); err != nil {
		return "", err
	}

	item, _ := c.GET(key)
	return resp.SerializeCacheItem(item), nil
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
	_, isExists := c.GET(key)
	return resp.SerializeBool(isExists)
}

func handleIncDec(key, operation string) (string, error) {
	if !c.EXISTS(key) {
		c.SET(key, 0)
	} else {
		data, _ := c.GET(key)
		if data.DataType != utils.Int {
			return "", errors.Err{Type: errors.TypeError}
		}
	}

	result := cache.DataItem{
		DataType: utils.Int,
	}
	switch operation {
	case resp.INCR:
		result.Value = []byte(strconv.Itoa(c.INCR(key)))
	case resp.DECR:
		result.Value = []byte(strconv.Itoa(c.DECR(key)))
	}

	if err := validators.ValidateExpiry(key); err != nil {
		return "", err
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
	_, keyExists := c.GET(key)
	if keyExists {
		c.DEL(key)
	}
	return resp.SerializeBool(keyExists)
}

func handleLpushRpush(key string, args []string, operation string) (string, error) {
	data, _ := c.GET(key)
	if data.DataType != utils.Array {
		return "", errors.Err{Type: errors.TypeError}
	}

	switch operation {
	case resp.LPUSH:
		c.LPUSH(key, args)
	case resp.RPUSH:
		c.RPUSH(key, args)
	}

	item := cache.DataItem{
		DataType: utils.Int,
	}

	vals := c.LRANGE(key, 0, -1)
	item.Value = []byte(strconv.Itoa(len(vals)))

	if err := validators.ValidateExpiry(key); err != nil {
		return "", err
	}
	return resp.SerializeCacheItem(item), nil
}

func handleLPUSH(key string, args []string) (string, error) {
	return handleLpushRpush(key, args, resp.LPUSH)
}

func handleRPUSH(key string, args []string) (string, error) {
	return handleLpushRpush(key, args, resp.RPUSH)
}

func handleLRANGE(key string, startIdx, endIdx int) (string, error) {
	data, _ := c.GET(key)
	if data.DataType != utils.Array {
		return "", errors.Err{Type: errors.TypeError}
	}

	if startIdx > endIdx {
		return "", errors.Err{Type: errors.IndexError}
	}

	vals := c.LRANGE(key, startIdx, endIdx)
	valsInBytes, _ := json.Marshal(vals)
	item := cache.DataItem{
		DataType: utils.Array,
		Value:    valsInBytes,
	}

	if err := validators.ValidateExpiry(key); err != nil {
		return "", err
	}
	return resp.SerializeCacheItem(item), nil
}

func handleListPop(key, popType string) (string, error) {
	data, _ := c.GET(key)
	if data.DataType != utils.Array {
		return "", errors.Err{Type: errors.TypeError}
	}

	var vals []string
	json.Unmarshal(data.Value, &vals)
	if len(vals) <= 0 {
		return "", errors.Err{Type: errors.EmptyList}
	}

	cacheItem := cache.DataItem{
		DataType: utils.String,
	}
	switch popType {
	case resp.LPOP:
		cacheItem.Value = []byte(c.LPOP(key))
	case resp.RPOP:
		cacheItem.Value = []byte(c.RPOP(key))
	}

	if err := validators.ValidateExpiry(key); err != nil {
		return "", err
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
	if !c.EXISTS(key) {
		return "", errors.Err{Type: errors.UndefinedKey}
	}

	c.EXPIRE(key, ttl)
	return "+OK\r\n", nil
}

func handleTTL(key string) (string, error) {
	data, _ := c.GET(key)
	if data.DataType != utils.Array {
		return "", errors.Err{Type: errors.TypeError}
	}

	cacheItem := cache.DataItem{
		DataType: utils.Int,
	}

	item, _ := c.GET(key)
	if item.ExpiryTime == nil {
		cacheItem.Value = []byte(strconv.Itoa(0))
	} else {
		if remainingTTL := int(time.Until(*item.ExpiryTime).Seconds()); remainingTTL > 0 {
			cacheItem.Value = []byte(strconv.Itoa(remainingTTL))
		} else {
			cacheItem.Value = []byte(strconv.Itoa(-1))
		}
	}
	return resp.SerializeCacheItem(cacheItem), nil
}

func handlePERSIST(key string) (string, error) {
	if c.EXISTS(key) {
		return "", errors.Err{Type: errors.UndefinedKey}
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
