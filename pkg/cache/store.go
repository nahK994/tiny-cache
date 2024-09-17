package cache

import (
	"fmt"
)

func InitCache() *Cache {
	return &Cache{
		info: make(map[string]interface{}),
	}
}

func (c *Cache) ReadCache(key string) interface{} {
	c.mu.RLock()         // Acquire read lock
	defer c.mu.RUnlock() // Release read lock
	return c.info[key]
}

func (c *Cache) WriteCache(key string, value interface{}) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock
	c.info[key] = value
}

func (c *Cache) IsKeyExist(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, isKeyExists := c.info[key]
	return isKeyExists
}

func (c *Cache) DeleteCache(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.info, key)
}

func (c *Cache) INCRCache(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, _ := c.info[key].(int)
	c.info[key] = val + 1
	return fmt.Sprintf(":%d\r\n", c.info[key])
}

func (c *Cache) DECRCache(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, _ := c.info[key].(int)

	c.info[key] = val - 1
	return fmt.Sprintf(":%d\r\n", c.info[key])
}

func (c *Cache) LPUSH(key string, value interface{}) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock
	c.info[key] = value
}

func (c *Cache) LRANGE(key string, startIdx, endIdx int) []string {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock

	vals, _ := c.info[key].([]string)
	if startIdx < 0 {
		startIdx = len(vals) + startIdx
	} else {
		startIdx--
	}
	if startIdx < 0 {
		startIdx = 0
	}

	if endIdx < 0 {
		endIdx = len(vals) + endIdx
	} else {
		endIdx--
	}
	if endIdx >= len(vals) {
		endIdx = len(vals) - 1
	}

	var ans []string
	for i := endIdx; i >= startIdx; i-- {
		ans = append(ans, vals[i])
	}
	return ans
}
