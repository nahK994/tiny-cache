package utils

import (
	"fmt"
	"strings"
)

type RESPCommands struct {
	SET      string
	GET      string
	LPUSH    string
	RPUSH    string
	LPOP     string
	RPOP     string
	EXPIRE   string
	DEL      string
	INCR     string
	DECR     string
	HSET     string
	HGET     string
	HMSET    string
	HMGET    string
	SMEMBERS string
	SADD     string
	SREM     string
	ZADD     string
	ZREM     string
	ZRANGE   string
	PING     string
	AUTH     string
	MULTI    string
	EXEC     string
	DISCARD  string
	WATCH    string
	UNWATCH  string
	MSET     string
	MGET     string
	LRANGE   string
	EXISTS   string
}

var respCommands = RESPCommands{
	SET:      "SET",
	GET:      "GET",
	LPUSH:    "LPUSH",
	RPUSH:    "RPUSH",
	LPOP:     "LPOP",
	RPOP:     "RPOP",
	EXPIRE:   "EXPIRE",
	DEL:      "DEL",
	INCR:     "INCR",
	DECR:     "DECR",
	HSET:     "HSET",
	HGET:     "HGET",
	HMSET:    "HMSET",
	HMGET:    "HMGET",
	SMEMBERS: "SMEMBERS",
	SADD:     "SADD",
	SREM:     "SREM",
	ZADD:     "ZADD",
	ZREM:     "ZREM",
	ZRANGE:   "ZRANGE",
	PING:     "PING",
	AUTH:     "AUTH",
	MULTI:    "MULTI",
	EXEC:     "EXEC",
	DISCARD:  "DISCARD",
	WATCH:    "WATCH",
	UNWATCH:  "UNWATCH",
	MSET:     "MSET",
	MGET:     "MGET",
	LRANGE:   "LRANGE",
	EXISTS:   "EXISTS",
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

var clientMessage string = fmt.Sprintf("Please use thses following commands:\n%s\n", strings.Join([]string{
	respCommands.PING, respCommands.EXISTS,
	respCommands.SET, respCommands.GET, respCommands.INCR, respCommands.DECR,
}, ", "))
