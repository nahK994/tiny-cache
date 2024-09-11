package cache

import "sync"

type Cache struct {
	info map[string]interface{}
	mu   sync.RWMutex
}
