package cache

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
	c.mu.RLock()
	defer c.mu.RUnlock()

	delete(c.info, key)
}
