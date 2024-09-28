package cache

import (
	"sync"
	"time"
)

type CacheData struct {
	DataType rune
	IntData  *int
	StrData  *string
	StrList  []string
}

type CacheItem struct {
	Value      CacheData
	ExpiryTime *time.Time
}

type Cache struct {
	items         map[string]CacheItem
	mu            sync.RWMutex
	SweepInterval int
}
