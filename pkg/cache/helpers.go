package cache

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/nahK994/TinyCache/pkg/utils"
)

func (c *Cache) evictLFU() {
	var lfuKey string
	minFreq := int(^uint(0) >> 1) // max int

	for key, item := range c.data {
		if item.Frequency < minFreq {
			minFreq = item.Frequency
			lfuKey = key
		}
	}

	delete(c.data, lfuKey)
}

// saveData stores a value in the cache
func (c *Cache) saveData(key string, value interface{}) {
	switch v := value.(type) {
	case int:
		c.saveInt(key, v, nil, 1)
	case string:
		if val, err := strconv.Atoi(v); err == nil {
			c.saveInt(key, val, nil, 1)
		} else {
			c.saveString(key, v, nil, 1)
		}
	}
}

// saveSting stores an string value
func (c *Cache) saveString(key string, value string, expiryTime *time.Time, frequency int) {
	bytes := []byte(value)
	c.data[key] = DataItem{
		DataType:   utils.String,
		Value:      bytes,
		ExpiryTime: expiryTime,
		Frequency:  frequency,
	}
}

// saveInt stores an integer value
func (c *Cache) saveInt(key string, value int, expiryTime *time.Time, frequency int) {
	bytes := []byte(strconv.Itoa(value))
	c.data[key] = DataItem{
		DataType:   utils.Int,
		Value:      bytes,
		ExpiryTime: expiryTime,
		Frequency:  frequency,
	}
}

func (c *Cache) IncrementFrequency(key string) error {
	item := c.data[key]
	item.Frequency++

	c.data[key] = item
	return nil
}

// saveList stores a list
func (c *Cache) saveList(key string, values []string, expiryTime *time.Time) {
	bytes, _ := json.Marshal(values)
	c.data[key] = DataItem{
		DataType:   utils.Array,
		Value:      bytes,
		ExpiryTime: expiryTime,
	}
}

// getList retrieves a list from cache
func (c *Cache) getList(key string) []string {
	item, exists := c.data[key]
	if !exists {
		return []string{}
	}

	var vals []string
	json.Unmarshal(item.Value, &vals)
	return vals
}

// reverseSlice reverses a slice of strings
func reverseSlice(s []string) []string {
	n := len(s)
	result := make([]string, n)
	for i := range s {
		result[n-i-1] = s[i]
	}
	return result
}

// processIdx normalizes index values for LRANGE
func processIdx(vals []string, idx int) int {
	if idx < 0 {
		idx = len(vals) + idx
	}
	if idx < 0 {
		return 0
	}
	if idx >= len(vals) {
		return len(vals) - 1
	}
	return idx
}
