package cache

import (
	"sync"
	"time"
)

type Data struct {
	Val        interface{}
	ExpiryTime time.Time
}

type Cache struct {
	Info map[string]Data
	mu   sync.RWMutex
}
