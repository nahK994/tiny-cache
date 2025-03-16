package cache

import (
	"sync"
	"time"
)

type DataItem struct {
	DataType   rune
	Value      []byte
	ExpiryTime *time.Time
}

type Cache struct {
	data          map[string]DataItem
	mu            sync.RWMutex
	SweepInterval int
}
