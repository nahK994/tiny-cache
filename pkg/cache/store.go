package cache

import (
	"strconv"
	"time"
)

// Init initializes the cache with an expiration sweeper
func NewCache(expirationSweepInterval, maxSize int) *Cache {
	cache := &Cache{
		data:          make(map[string]DataItem),
		SweepInterval: expirationSweepInterval,
		MaxSize:       maxSize,
	}
	go cache.activeExpiration()
	return cache
}

// GET retrieves an item from the cache
func (c *Cache) GET(key string) (DataItem, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.data[key]
	return item, exists
}

// SET stores an item in the cache
func (c *Cache) SET(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.saveData(key, value)
	if len(c.data) >= c.MaxSize {
		c.evictLFU()
	}
}

// EXISTS checks if a key exists in the cache
func (c *Cache) EXISTS(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.data[key]
	return exists
}

// DEL removes a key from the cache
func (c *Cache) DEL(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// INCR increments an integer value
func (c *Cache) INCR(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := c.data[key]
	intVal, _ := strconv.Atoi(string(item.Value))
	intVal++

	c.saveInt(key, intVal, item.ExpiryTime, item.Frequency)
	return intVal
}

// DECR decrements an integer value
func (c *Cache) DECR(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := c.data[key]
	intVal, _ := strconv.Atoi(string(item.Value))
	intVal--

	c.saveInt(key, intVal, item.ExpiryTime, item.Frequency)
	return intVal
}

// LPUSH adds values to the left of a list
func (c *Cache) LPUSH(key string, values []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	oldItem := c.data[key]
	vals := append(reverseSlice(values), getList(oldItem.Value)...)

	c.data[key] = c.createListItem(vals, oldItem.ExpiryTime, oldItem.Frequency)
}

// RPUSH adds values to the right of a list
func (c *Cache) RPUSH(key string, values []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	oldItem := c.data[key]
	vals := append(getList(oldItem.Value), values...)

	c.data[key] = c.createListItem(vals, oldItem.ExpiryTime, oldItem.Frequency)
}

// LRANGE retrieves a range of values from a list
func (c *Cache) LRANGE(key string, startIdx, endIdx int) []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	vals := getList(c.data[key].Value)
	if len(vals) == 0 {
		return []string{}
	}

	startIdx = processIdx(vals, startIdx)
	endIdx = processIdx(vals, endIdx)

	return vals[startIdx : endIdx+1]
}

// LPOP removes and returns the first element from a list
func (c *Cache) LPOP(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	vals := getList(c.data[key].Value)

	oldItem := c.data[key]
	c.data[key] = c.createListItem(vals[1:], oldItem.ExpiryTime, oldItem.Frequency)
	return vals[0]
}

// RPOP removes and returns the last element from a list
func (c *Cache) RPOP(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	vals := getList(c.data[key].Value)

	oldItem := c.data[key]
	c.data[key] = c.createListItem(vals[:len(vals)-1], oldItem.ExpiryTime, oldItem.Frequency)
	return vals[len(vals)-1]
}

// FLUSHALL removes all keys from the cache
func (c *Cache) FLUSHALL() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]DataItem)
}

// EXPIRE sets a time-to-live (TTL) for a key
func (c *Cache) EXPIRE(key string, ttl int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, exists := c.data[key]
	if !exists {
		return
	}

	if ttl > 0 {
		expiry := time.Now().Add(time.Duration(ttl) * time.Second)
		item.ExpiryTime = &expiry
	} else {
		item.ExpiryTime = nil
	}

	c.data[key] = item
}

// activeExpiration removes expired keys periodically
func (c *Cache) activeExpiration() {
	for {
		time.Sleep(time.Duration(c.SweepInterval) * time.Second)

		c.mu.Lock()
		for key, item := range c.data {
			if item.ExpiryTime != nil && time.Now().After(*item.ExpiryTime) {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
