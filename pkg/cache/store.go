package cache

import (
	"time"
)

func Init(expirationSweepInterval int) *Cache {
	cache := &Cache{
		items:         make(map[string]CacheItem),
		SweepInterval: expirationSweepInterval,
	}
	go cache.activeExpiration()
	return cache
}

func (c *Cache) GET(key string) CacheItem {
	c.mu.RLock()         // Acquire read lock
	defer c.mu.RUnlock() // Release read lock
	return c.items[key]
}

func (c *Cache) SET(key string, value interface{}) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock
	c.saveData(key, value)
}

func (c *Cache) EXISTS(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, exists := c.items[key]
	return exists
}

func (c *Cache) DEL(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

func (c *Cache) INCR(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	val := c.items[key].Value.IntData
	*val++
	c.saveInt(key, val)
	return *val
}

func (c *Cache) DECR(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	val := c.items[key].Value.IntData
	*val--
	c.saveInt(key, val)
	return *val
}

func (c *Cache) LPUSH(key string, values []string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock

	oldData := c.items[key].Value.StrList
	vals := make([]string, len(values)+len(oldData))
	for i := 0; i < len(values); i++ {
		vals[i] = values[len(values)-1-i]
	}
	copy(vals[len(values):], oldData)
	oldData = nil
	c.saveList(key, vals)
}

func (c *Cache) RPUSH(key string, values []string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock

	oldData := c.items[key].Value.StrList
	vals := make([]string, len(values)+len(oldData))
	copy(vals[0:], oldData)
	copy(vals[len(oldData):], values)
	oldData = nil
	c.saveList(key, vals)
}

func (c *Cache) LRANGE(key string, startIdx, endIdx int) []string {
	c.mu.RLock()         // Acquire write lock
	defer c.mu.RUnlock() // Release write lock

	vals := c.items[key].Value.StrList
	startIdx = processIdx(vals, startIdx)
	endIdx = processIdx(vals, endIdx)
	if len(vals) > 0 && startIdx <= endIdx {
		return vals[startIdx : endIdx+1]
	}
	return nil
}

func (c *Cache) LPOP(key string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock

	vals := c.items[key].Value.StrList
	newVals := make([]string, len(vals)-1)
	copy(newVals, vals[1:])

	c.saveList(key, newVals)
}

func (c *Cache) RPOP(key string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock

	vals := c.items[key].Value.StrList
	newVals := make([]string, len(vals)-1)
	copy(newVals, vals[0:len(vals)-1])

	c.saveList(key, newVals)
}

func (c *Cache) FLUSHALL() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k := range c.items {
		delete(c.items, k)
	}
}

func (c *Cache) EXPIRE(key string, ttl int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val := c.items[key]

	if ttl != 0 {
		tt := time.Now().Add(time.Duration(ttl) * time.Second)
		val.ExpiryTime = &tt
	} else {
		val.ExpiryTime = nil
	}
	c.items[key] = val
}

func (c *Cache) activeExpiration() {
	for {
		time.Sleep(time.Duration(c.SweepInterval * int(time.Second)))
		c.mu.Lock()
		for key, item := range c.items {
			if item.ExpiryTime != nil && time.Now().After(*item.ExpiryTime) {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
