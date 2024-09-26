package utils

import (
	"github.com/nahK994/TinyCache/pkg/cache"
	"github.com/nahK994/TinyCache/pkg/config"
	"github.com/nahK994/TinyCache/pkg/errors"
)

var errType errors.ErrTypes = errors.GetErrorTypes()
var c *cache.Cache = config.App.Cache

type RESPCommands struct {
	SET      string
	GET      string
	LPUSH    string
	RPUSH    string
	LPOP     string
	RPOP     string
	DEL      string
	INCR     string
	DECR     string
	PING     string
	LRANGE   string
	EXISTS   string
	FLUSHALL string
	EXPIRE   string
	TTL      string
	PERSIST  string
}

var respCmds = RESPCommands{
	SET:      "SET",
	GET:      "GET",
	LPUSH:    "LPUSH",
	RPUSH:    "RPUSH",
	LPOP:     "LPOP",
	RPOP:     "RPOP",
	DEL:      "DEL",
	INCR:     "INCR",
	DECR:     "DECR",
	PING:     "PING",
	LRANGE:   "LRANGE",
	EXISTS:   "EXISTS",
	FLUSHALL: "FLUSHALL",
	EXPIRE:   "EXPIRE",
	TTL:      "TTL",
	PERSIST:  "PERSIST",
}

type ReplyType struct {
	Int    rune
	Bulk   rune
	Error  rune
	Array  rune
	Status rune
}

var replyType ReplyType = ReplyType{
	Int:    ':',
	Bulk:   '$',
	Array:  '*',
	Error:  '-',
	Status: '+',
}
