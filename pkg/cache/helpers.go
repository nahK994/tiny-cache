package cache

import (
	"strconv"
)

func (c *Cache) saveData(key string, val interface{}) {
	switch v := val.(type) {
	case int:
		c.saveInt(key, &v)
	case string:
		if val, err := strconv.Atoi(v); err == nil {
			c.saveInt(key, &val)
		} else {
			c.saveString(key, &v)
		}
	case []string:
		c.saveList(key, v)
	}

}

func (c *Cache) saveString(key string, val *string) {
	c.items[key] = CacheItem{
		Value: CacheData{
			StrData:  val,
			DataType: String,
		},
		ExpiryTime: c.items[key].ExpiryTime,
	}
}

func (c *Cache) saveInt(key string, val *int) {
	c.items[key] = CacheItem{
		Value: CacheData{
			IntData:  val,
			DataType: Int,
		},
		ExpiryTime: c.items[key].ExpiryTime,
	}
}

func (c *Cache) saveList(key string, val []string) {
	c.items[key] = CacheItem{
		Value: CacheData{
			StrList:  val,
			DataType: Array,
		},
		ExpiryTime: c.items[key].ExpiryTime,
	}
}

func processIdx(vals []string, idx int) int {
	if idx > len(vals) {
		idx = len(vals) - 1
	} else if idx < 0 {
		if -1*len(vals) > idx {
			idx = 0
		} else {
			idx = len(vals) + idx
		}
	}

	return idx
}
