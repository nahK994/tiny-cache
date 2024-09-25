package cache

import (
	"strconv"
	"time"
)

func InitCache(expirationSweepInterval int) *Cache {
	cache := &Cache{
		Info:                    make(map[string]Data),
		ExpirationSweepInterval: expirationSweepInterval,
	}
	go cache.activeExpiration()
	return cache
}

func (c *Cache) GET(key string) Data {
	c.mu.RLock()         // Acquire read lock
	defer c.mu.RUnlock() // Release read lock
	return c.Info[key]
}

func (c *Cache) SET(key string, value interface{}) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock
	str, ok_str := value.(string)
	if ok_str {
		num, err := strconv.Atoi(str)
		if err == nil {
			c.Info[key] = Data{
				Val: num,
			}
		} else {
			c.Info[key] = Data{
				Val: str,
			}
		}
	} else {
		c.Info[key] = Data{
			Val: value,
		}
	}
}

func (c *Cache) EXISTS(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, exists := c.Info[key]
	return exists
}

func (c *Cache) DEL(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.Info, key)
}

func (c *Cache) INCR(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, _ := c.Info[key].Val.(int)
	c.Info[key] = Data{
		Val: val + 1,
	}
	return val + 1
}

func (c *Cache) DECR(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, _ := c.Info[key].Val.(int)
	c.Info[key] = Data{
		Val: val - 1,
	}
	return val - 1
}

func (c *Cache) LPUSH(key string, values []string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock

	oldData, _ := c.Info[key].Val.([]string)
	vals := make([]string, len(values)+len(oldData))
	for i := 0; i < len(values); i++ {
		vals[i] = values[len(values)-1-i]
	}
	copy(vals[len(values):], oldData)
	oldData = nil
	c.Info[key] = Data{
		Val: vals,
	}
}

func (c *Cache) RPUSH(key string, values []string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock

	oldData, _ := c.Info[key].Val.([]string)
	vals := make([]string, len(values)+len(oldData))
	copy(vals[0:], oldData)
	copy(vals[len(oldData):], values)
	oldData = nil
	c.Info[key] = Data{
		Val: vals,
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

func (c *Cache) LRANGE(key string, startIdx, endIdx int) []string {
	c.mu.RLock()         // Acquire write lock
	defer c.mu.RUnlock() // Release write lock

	vals, _ := c.Info[key].Val.([]string)
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

	vals, _ := c.Info[key].Val.([]string)

	newVals := make([]string, len(vals)-1)
	copy(newVals, vals[1:])

	c.Info[key] = Data{
		Val: newVals,
	}
	vals = nil
}

func (c *Cache) RPOP(key string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock

	vals, _ := c.Info[key].Val.([]string)

	newVals := make([]string, len(vals)-1)
	copy(newVals, vals[0:len(vals)-1])

	c.Info[key] = Data{
		Val: newVals,
	}
	vals = nil
}

func (c *Cache) FLUSHALL() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k := range c.Info {
		delete(c.Info, k)
	}
}

func (c *Cache) EXPIRE(key string, ttl int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val := c.Info[key]

	val.ExpiryTime = time.Now().Add(time.Second * time.Duration(ttl))
	c.Info[key] = val
}

func (c *Cache) activeExpiration() {
	for {
		time.Sleep(time.Duration(c.ExpirationSweepInterval * int(time.Second)))
		c.mu.Lock()
		for key, item := range c.Info {
			if time.Now().After(item.ExpiryTime) {
				delete(c.Info, key)
			}
		}
		c.mu.Unlock()
	}
}
