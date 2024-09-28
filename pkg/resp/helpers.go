package resp

import (
	"fmt"

	"github.com/nahK994/TinyCache/pkg/cache"
)

func processArray(segments []string) string {
	response := fmt.Sprintf("*%d\r\n", len(segments))
	for _, v := range segments {
		response += fmt.Sprintf("$%d\r\n%s\r\n", len(v), v)
	}
	return response
}

func SerializeBool(arg bool) string {
	if arg {
		return ":1\r\n"
	} else {
		return ":0\r\n"
	}
}

func SerializeCacheItem(item cache.CacheData) string {
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
			return processArray(item.StrList)
		}
	}
	return "$-1\r\n"
}
