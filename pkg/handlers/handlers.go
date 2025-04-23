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

func (h *Handler) handleGET(key string) (*cache.DataItem, error) {
	if !h.cache.EXISTS(key) {
		return nil, errors.Err{Type: errors.UndefinedKey}
	}

	if err := h.validateExpiry(key); err != nil {
		return nil, err
	}

	item, _ := h.cache.GET(key)
	return item, nil
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

func (h *Handler) handleEXISTS(key string) bool {
	_, isExists := h.cache.GET(key)
	return isExists
}

func (h *Handler) handleIncDec(key, operation string) (*cache.DataItem, error) {
	if !h.cache.EXISTS(key) {
		h.cache.SET(key, 0)
	} else {
		data, _ := h.cache.GET(key)
		if data.DataType != utils.Int {
			return nil, errors.Err{Type: errors.TypeError}
		}
	}

	var val []byte
	switch operation {
	case resp.INCR:
		val = []byte(strconv.Itoa(h.cache.INCR(key)))
	case resp.DECR:
		val = []byte(strconv.Itoa(h.cache.DECR(key)))
	}

	if err := h.validateExpiry(key); err != nil {
		return nil, err
	}

	return &cache.DataItem{
		DataType: utils.Int,
		Value:    val,
	}, nil
}

func (h *Handler) handleDEL(key string) bool {
	_, keyExists := h.cache.GET(key)
	if keyExists {
		h.cache.DEL(key)
	}
	return keyExists
}

func (h *Handler) handleLpushRpush(key string, args []string, operation string) (*cache.DataItem, error) {
	data, isExists := h.cache.GET(key)

	if isExists && data.DataType != utils.Array {
		return nil, errors.Err{Type: errors.TypeError}
	}

	switch operation {
	case resp.LPUSH:
		h.cache.LPUSH(key, args)
	case resp.RPUSH:
		h.cache.RPUSH(key, args)
	}

	vals := h.cache.LRANGE(key, 0, -1)
	if err := h.validateExpiry(key); err != nil {
		return nil, err
	}

	return &cache.DataItem{
		DataType: utils.Int,
		Value:    []byte(strconv.Itoa(len(vals))),
	}, nil
}

func (h *Handler) handleLRANGE(key string, startIdx, endIdx int) (*cache.DataItem, error) {
	data, isExists := h.cache.GET(key)
	if !isExists {
		return nil, errors.Err{Type: errors.UndefinedKey}
	}

	if data.DataType != utils.Array {
		return nil, errors.Err{Type: errors.TypeError}
	}

	if startIdx > endIdx {
		return nil, errors.Err{Type: errors.IndexError}
	}

	vals := h.cache.LRANGE(key, startIdx, endIdx)
	valsInBytes, _ := json.Marshal(vals)

	if err := h.validateExpiry(key); err != nil {
		return nil, err
	}

	return &cache.DataItem{
		DataType: utils.Array,
		Value:    valsInBytes,
	}, nil
}

func (h *Handler) handleListPop(key, popType string) (*cache.DataItem, error) {
	data, _ := h.cache.GET(key)
	if data.DataType != utils.Array {
		return nil, errors.Err{Type: errors.TypeError}
	}

	var vals []string
	json.Unmarshal(data.Value, &vals)
	if len(vals) <= 0 {
		return nil, errors.Err{Type: errors.EmptyList}
	}

	var val []byte
	switch popType {
	case resp.LPOP:
		val = []byte(h.cache.LPOP(key))
	case resp.RPOP:
		val = []byte(h.cache.RPOP(key))
	}

	if err := h.validateExpiry(key); err != nil {
		return nil, err
	}
	return &cache.DataItem{
		DataType: utils.String,
		Value:    val,
	}, nil
}

func (h *Handler) handleEXPIRE(key string, ttl int) error {
	if !h.cache.EXISTS(key) {
		return errors.Err{Type: errors.UndefinedKey}
	}

	h.cache.EXPIRE(key, ttl)
	return nil
}

func (h *Handler) handleTTL(key string) (*cache.DataItem, error) {
	data, isExists := h.cache.GET(key)
	if !isExists {
		return nil, errors.Err{Type: errors.UndefinedKey}
	}

	var val []byte
	if data.ExpiryTime == nil {
		val = []byte(strconv.Itoa(0))
	} else {
		if remainingTTL := int(time.Until(*data.ExpiryTime).Seconds()); remainingTTL > 0 {
			val = []byte(strconv.Itoa(remainingTTL))
		} else {
			h.cache.DEL(key)
			return nil, errors.Err{Type: errors.ExpiredKey}
		}
	}

	return &cache.DataItem{
		DataType: utils.Int,
		Value:    val,
	}, nil
}

func (h *Handler) handlePERSIST(key string) error {
	if !h.cache.EXISTS(key) {
		return errors.Err{Type: errors.UndefinedKey}
	}

	h.cache.EXPIRE(key, 0)
	return nil
}

func (h *Handler) HandleCommand(serializedRawCmd string) (string, error) {
	cmdSegments, _ := resp.Deserializer(serializedRawCmd).([]string)

	cmd := cmdSegments[0]
	args := cmdSegments[1:]
	var key string
	if len(args) > 0 {
		key = args[0]
	}

	switch strings.ToUpper(cmd) {
	case resp.GET:
		response, err := h.handleGET(key)
		if err != nil {
			return "", err
		}
		h.cache.IncrementFrequency(key)
		return resp.SerializeCacheItem(response), nil
	case resp.SET:
		return h.handleSET(key, args[1:])
	case resp.EXISTS:
		isExists := h.handleEXISTS(key)
		if isExists {
			h.cache.IncrementFrequency(key)
		}
		return resp.SerializeBool(isExists), nil
	case resp.DEL:
		isExists := h.handleDEL(key)
		return resp.SerializeBool(isExists), nil
	case resp.PING:
		return "+PONG\r\n", nil
	case resp.LPUSH:
		response, err := h.handleLpushRpush(key, args[1:], resp.LPUSH)
		if err != nil {
			return "", err
		}
		h.cache.IncrementFrequency(key)
		return resp.SerializeCacheItem(response), nil
	case resp.RPUSH:
		response, err := h.handleLpushRpush(key, args[1:], resp.RPUSH)
		if err != nil {
			return "", err
		}
		h.cache.IncrementFrequency(key)
		return resp.SerializeCacheItem(response), nil
	case resp.EXPIRE:
		ttl, _ := strconv.Atoi(args[1])
		err := h.handleEXPIRE(key, ttl)
		if err != nil {
			return "", err
		}
		h.cache.IncrementFrequency(key)
		return "+OK\r\n", nil
	case resp.LRANGE:
		startIdx, _ := strconv.Atoi(args[1])
		endIdx, _ := strconv.Atoi(args[2])
		response, err := h.handleLRANGE(key, startIdx, endIdx)
		if err != nil {
			return "", err
		}
		h.cache.IncrementFrequency(key)
		return resp.SerializeCacheItem(response), nil
	case resp.LPOP:
		response, err := h.handleListPop(key, resp.LPOP)
		if err != nil {
			return "", err
		}
		h.cache.IncrementFrequency(key)
		return resp.SerializeCacheItem(response), nil
	case resp.RPOP:
		response, err := h.handleListPop(key, resp.RPOP)
		if err != nil {
			return "", err
		}
		h.cache.IncrementFrequency(key)
		return resp.SerializeCacheItem(response), nil
	case resp.INCR:
		response, err := h.handleIncDec(key, resp.INCR)
		if err != nil {
			return "", err
		}
		h.cache.IncrementFrequency(key)
		return resp.SerializeCacheItem(response), nil
	case resp.DECR:
		response, err := h.handleIncDec(key, resp.DECR)
		if err != nil {
			return "", err
		}
		h.cache.IncrementFrequency(key)
		return resp.SerializeCacheItem(response), nil
	case resp.FLUSHALL:
		return h.handleFLUSHALL(), nil
	case resp.TTL:
		response, err := h.handleTTL(key)
		if err != nil {
			return "", err
		}
		h.cache.IncrementFrequency(key)
		return resp.SerializeCacheItem(response), nil
	case resp.PERSIST:
		err := h.handlePERSIST(key)
		if err != nil {
			return "", err
		}
		h.cache.IncrementFrequency(key)
		return "+OK\r\n", nil
	default:
		return "", errors.Err{Type: errors.UnknownCommand}
	}
}
