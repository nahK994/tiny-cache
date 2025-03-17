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
)

type Handler struct {
	cache *cache.Cache
}

func NewHandler(c *cache.Cache) *Handler {
	return &Handler{
		cache: c,
	}
}

func (h *Handler) validateExpiry(key string) error {
	item, _ := h.cache.GET(key)
	if item.ExpiryTime != nil && time.Now().After(*item.ExpiryTime) {
		h.cache.DEL(key)
		return errors.Err{Type: errors.ExpiredKey}
	}
	return nil
}

func (h *Handler) handleGET(key string) (string, error) {
	if !h.cache.EXISTS(key) {
		return "", errors.Err{Type: errors.UndefinedKey}
	}

	if err := h.validateExpiry(key); err != nil {
		return "", err
	}

	item, _ := h.cache.GET(key)
	return resp.SerializeCacheItem(item), nil
}

func (h *Handler) handleSET(key string, args []string) (string, error) {
	if len(args) > 1 {
		h.cache.SET(key, args[0])
		ttl, _ := strconv.Atoi(args[1])
		h.cache.EXPIRE(key, ttl)
	} else {
		h.cache.SET(key, args[0])
	}
	return "+OK\r\n", nil
}

func (h *Handler) handleFLUSHALL() string {
	if config.App.IsAsyncFlush {
		config.App.FlushCh <- 1
	} else {
		h.cache.FLUSHALL()
	}
	return "+OK\r\n"
}

func (h *Handler) handleEXISTS(key string) string {
	_, isExists := h.cache.GET(key)
	return resp.SerializeBool(isExists)
}

func (h *Handler) handleIncDec(key, operation string) (string, error) {
	if !h.cache.EXISTS(key) {
		h.cache.SET(key, 0)
	} else {
		data, _ := h.cache.GET(key)
		if data.DataType != utils.Int {
			return "", errors.Err{Type: errors.TypeError}
		}
	}

	var val []byte
	switch operation {
	case resp.INCR:
		val = []byte(strconv.Itoa(h.cache.INCR(key)))
	case resp.DECR:
		val = []byte(strconv.Itoa(h.cache.DECR(key)))
	}

	result := cache.DataItem{
		DataType: utils.Int,
		Value:    val,
	}

	if err := h.validateExpiry(key); err != nil {
		return "", err
	}
	return resp.SerializeCacheItem(result), nil
}

func (h *Handler) handleINCR(key string) (string, error) {
	return h.handleIncDec(key, resp.INCR)
}

func (h *Handler) handleDECR(key string) (string, error) {
	return h.handleIncDec(key, resp.DECR)
}

func (h *Handler) handleDEL(key string) string {
	_, keyExists := h.cache.GET(key)
	if keyExists {
		h.cache.DEL(key)
	}
	return resp.SerializeBool(keyExists)
}

func (h *Handler) handleLpushRpush(key string, args []string, operation string) (string, error) {
	data, isExists := h.cache.GET(key)

	if isExists && data.DataType != utils.Array {
		return "", errors.Err{Type: errors.TypeError}
	}

	switch operation {
	case resp.LPUSH:
		h.cache.LPUSH(key, args)
	case resp.RPUSH:
		h.cache.RPUSH(key, args)
	}

	vals := h.cache.LRANGE(key, 0, -1)
	item := cache.DataItem{
		DataType: utils.Int,
		Value:    []byte(strconv.Itoa(len(vals))),
	}

	if err := h.validateExpiry(key); err != nil {
		return "", err
	}
	return resp.SerializeCacheItem(item), nil
}

func (h *Handler) handleLPUSH(key string, args []string) (string, error) {
	return h.handleLpushRpush(key, args, resp.LPUSH)
}

func (h *Handler) handleRPUSH(key string, args []string) (string, error) {
	return h.handleLpushRpush(key, args, resp.RPUSH)
}

func (h *Handler) handleLRANGE(key string, startIdx, endIdx int) (string, error) {
	data, isExists := h.cache.GET(key)
	if !isExists {
		return "", errors.Err{Type: errors.UndefinedKey}
	}

	if data.DataType != utils.Array {
		return "", errors.Err{Type: errors.TypeError}
	}

	if startIdx > endIdx {
		return "", errors.Err{Type: errors.IndexError}
	}

	vals := h.cache.LRANGE(key, startIdx, endIdx)
	valsInBytes, _ := json.Marshal(vals)
	item := cache.DataItem{
		DataType: utils.Array,
		Value:    valsInBytes,
	}

	if err := h.validateExpiry(key); err != nil {
		return "", err
	}
	return resp.SerializeCacheItem(item), nil
}

func (h *Handler) handleListPop(key, popType string) (string, error) {
	data, _ := h.cache.GET(key)
	if data.DataType != utils.Array {
		return "", errors.Err{Type: errors.TypeError}
	}

	var vals []string
	json.Unmarshal(data.Value, &vals)
	if len(vals) <= 0 {
		return "", errors.Err{Type: errors.EmptyList}
	}

	var val []byte
	switch popType {
	case resp.LPOP:
		val = []byte(h.cache.LPOP(key))
	case resp.RPOP:
		val = []byte(h.cache.RPOP(key))
	}

	cacheItem := cache.DataItem{
		DataType: utils.String,
		Value:    val,
	}

	if err := h.validateExpiry(key); err != nil {
		return "", err
	}
	return resp.SerializeCacheItem(cacheItem), nil
}

func (h *Handler) handleLPOP(key string) (string, error) {
	return h.handleListPop(key, resp.LPOP)
}

func (h *Handler) handleRPOP(key string) (string, error) {
	return h.handleListPop(key, resp.RPOP)
}

func (h *Handler) handleEXPIRE(key string, ttl int) (string, error) {
	if !h.cache.EXISTS(key) {
		return "", errors.Err{Type: errors.UndefinedKey}
	}

	h.cache.EXPIRE(key, ttl)
	return "+OK\r\n", nil
}

func (h *Handler) handleTTL(key string) (string, error) {
	data, isExists := h.cache.GET(key)
	if !isExists {
		return "", errors.Err{Type: errors.UndefinedKey}
	}

	var val []byte
	if data.ExpiryTime == nil {
		val = []byte(strconv.Itoa(0))
	} else {
		if remainingTTL := int(time.Until(*data.ExpiryTime).Seconds()); remainingTTL > 0 {
			val = []byte(strconv.Itoa(remainingTTL))
		} else {
			h.cache.DEL(key)
			return "", errors.Err{Type: errors.ExpiredKey}
		}
	}

	cacheItem := cache.DataItem{
		DataType: utils.Int,
		Value:    val,
	}
	return resp.SerializeCacheItem(cacheItem), nil
}

func (h *Handler) handlePERSIST(key string) (string, error) {
	if !h.cache.EXISTS(key) {
		return "", errors.Err{Type: errors.UndefinedKey}
	}

	h.cache.EXPIRE(key, 0)
	return "+OK\r\n", nil
}

func (h *Handler) HandleCommand(serializedRawCmd string) (string, error) {
	cmdSegments, _ := resp.Deserializer(serializedRawCmd).([]string)

	cmd := cmdSegments[0]
	args := cmdSegments[1:]

	switch strings.ToUpper(cmd) {
	case resp.GET:
		return h.handleGET(args[0])
	case resp.SET:
		return h.handleSET(args[0], args[1:])
	case resp.EXISTS:
		return h.handleEXISTS(args[0]), nil
	case resp.DEL:
		return h.handleDEL(args[0]), nil
	case resp.PING:
		return "+PONG\r\n", nil
	case resp.LPUSH:
		return h.handleLPUSH(args[0], args[1:])
	case resp.RPUSH:
		return h.handleRPUSH(args[0], args[1:])
	case resp.EXPIRE:
		ttl, _ := strconv.Atoi(args[1])
		return h.handleEXPIRE(args[0], ttl)
	case resp.LRANGE:
		startIdx, _ := strconv.Atoi(args[1])
		endIdx, _ := strconv.Atoi(args[2])
		return h.handleLRANGE(args[0], startIdx, endIdx)
	case resp.LPOP:
		return h.handleLPOP(args[0])
	case resp.RPOP:
		return h.handleRPOP(args[0])
	case resp.INCR:
		return h.handleINCR(args[0])
	case resp.DECR:
		return h.handleDECR(args[0])
	case resp.FLUSHALL:
		return h.handleFLUSHALL(), nil
	case resp.TTL:
		return h.handleTTL(args[0])
	case resp.PERSIST:
		return h.handlePERSIST(args[0])
	default:
		return "", errors.Err{Type: errors.UnknownCommand}
	}
}
