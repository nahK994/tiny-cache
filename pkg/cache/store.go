package cache

import (
	"strconv"
	"time"
)

// Init initializes the cache with an expiration sweeper
func NewCache(expirationSweepInterval, maxSize int) *Cache {
	cache := &Cache{
		data:          make(map[string]*DataItem),
		SweepInterval: expirationSweepInterval,
		MaxSize:       maxSize,
	}
	go cache.activeExpiration()
	return cache
}

// GET retrieves an item from the cache
func (c *Cache) GET(key string) (*DataItem, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.data[key]
	return item, exists
}

// SET stores an item in the cache
func (c *Cache) SET(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var item *DataItem
	switch v := value.(type) {
	case int:
		item = createIntItem(v, nil, 0)
	case string:
		if val, err := strconv.Atoi(v); err == nil {
			item = createIntItem(val, nil, 0)
		} else {
			item = createStringItem(v, nil, 0)
		}
	}

	c.evictLFU()
	c.data[key] = item
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

func (c *Cache) IncrDecr(key string, delta int) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := c.data[key]
	intVal, _ := strconv.Atoi(string(item.Value))
	intVal += delta

	c.data[key] = createIntItem(intVal, item.ExpiryTime, item.Frequency)
	return intVal
}

func (c *Cache) LPUSH(key string, values []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	oldItem := c.data[key]
	var vals []string
	var expiry *time.Time
	var freq int

	if oldItem != nil {
		vals = append(reverseSlice(values), getList(oldItem.Value)...)
		expiry = oldItem.ExpiryTime
		freq = oldItem.Frequency
	} else {
		vals = reverseSlice(values)
		expiry = nil
		freq = 0
	}

	c.data[key] = createListItem(vals, expiry, freq)
}

func (c *Cache) RPUSH(key string, values []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	oldItem := c.data[key]
	var vals []string
	var expiry *time.Time
	var freq int

	if oldItem != nil {
		vals = append(getList(oldItem.Value), values...)
		expiry = oldItem.ExpiryTime
		freq = oldItem.Frequency
	} else {
		vals = values
		expiry = nil
		freq = 0
	}

	c.data[key] = createListItem(vals, expiry, freq)
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
	c.data[key] = createListItem(vals[1:], oldItem.ExpiryTime, oldItem.Frequency)
	return vals[0]
}

// RPOP removes and returns the last element from a list
func (c *Cache) RPOP(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	vals := getList(c.data[key].Value)

	oldItem := c.data[key]
	c.data[key] = createListItem(vals[:len(vals)-1], oldItem.ExpiryTime, oldItem.Frequency)
	return vals[len(vals)-1]
}

// FLUSHALL removes all keys from the cache
func (c *Cache) FLUSHALL() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]*DataItem)
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

func (c *Cache) IncrementFrequency(key string) error {
	item := c.data[key]
	(*item).Frequency++

	c.data[key] = item
	return nil
}

func (c *Cache) evictLFU() {
	if len(c.data) < c.MaxSize {
		return
	}
	var evictKey string
	minFreq := int(^uint(0) >> 1) // max int

	for key, item := range c.data {
		if item.Frequency < minFreq {
			minFreq = item.Frequency
			evictKey = key
		}
	}

	delete(c.data, evictKey)
}
