package cache

import (
	"sync"
	"time"
)

type data struct {
	val        interface{}
	expiryTime time.Time
}

type Cache struct {
	info map[string]data
	mu   sync.RWMutex
}
