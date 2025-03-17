package config

type app struct {
	Host         string
	Port         int
	IsAsyncFlush bool
	FlushCh      chan int
}

var App app = app{
	Host:         "127.0.0.1",
	Port:         8888,
	IsAsyncFlush: true,
	FlushCh:      make(chan int),
}
