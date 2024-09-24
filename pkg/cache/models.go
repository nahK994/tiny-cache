package cache

import (
	"sync"
	"time"
)

type Data struct {
	val        interface{}
	expiryTime time.Time
}

type Cache struct {
	info map[string]Data
	mu   sync.RWMutex
}
