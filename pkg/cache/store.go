package cache

import (
	"fmt"

	"github.com/nahK994/TinyCache/pkg/errors"
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

func (c *Cache) WriteCache(key string, value interface{}) error {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock
	c.info[key] = value
	return nil
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

func (c *Cache) INCRCache(key string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.info[key].(int)
	if !ok {
		return "", errors.Err{Msg: "-ERR value aren't available for INCR\r\n", File: "handlers/handlers.go", Line: 49}
	}

	c.info[key] = val + 1
	return fmt.Sprintf(":%d\r\n", c.info[key]), nil
}

func (c *Cache) DECRCache(key string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.info[key].(int)
	if !ok {
		return "", errors.Err{Msg: "-ERR value aren't available for DECR\r\n", File: "handlers/handlers.go", Line: 62}
	}

	c.info[key] = val - 1
	return fmt.Sprintf(":%d\r\n", c.info[key]), nil
}
