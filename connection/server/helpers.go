package server

import (
	"fmt"

	"github.com/nahK994/TinyCache/pkg/cache"
)

func IsKeyExists(key string) bool {
	_ = validateExpiry(key)
	return c.EXISTS(key)
}

func processCacheItem(item cache.CacheData) string {
	switch item.DataType {
	case cache.Int:
		if item.IntData != nil {
			return fmt.Sprintf(":%d\r\n", *item.IntData)
		}
	case cache.String:
		if item.StrData != nil {
			return fmt.Sprintf("$%d\r\n%s\r\n", len(*item.StrData), *item.StrData)
		}
	case cache.Array:
		if item.StrList != nil {
			response := fmt.Sprintf("*%d\r\n", len(item.StrList))
			for _, v := range item.StrList {
				response += fmt.Sprintf("$%d\r\n%s\r\n", len(v), v)
			}
			return response
		}
	}
	return "$-1\r\n"
}
