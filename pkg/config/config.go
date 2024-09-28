package config

import "github.com/nahK994/TinyCache/pkg/cache"

type app struct {
	Host         string
	Port         int
	IsAsyncFlush bool
	FlushCh      chan int
	Cache        *cache.Cache
}

var App app = app{
	Host:         "127.0.0.1",
	Port:         8888,
	IsAsyncFlush: true,
	FlushCh:      make(chan int),
	Cache:        cache.Init(60),
}
