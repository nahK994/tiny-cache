package resp

import (
	"github.com/nahK994/TinyCache/pkg/cache"
	"github.com/nahK994/TinyCache/pkg/shared"
)

type deserializeTestCase struct {
	input  string
	output []string
}

type serializeTestCase struct {
	input  string
	output string
}

type serializeCacheDataTestCase struct {
	input  cache.CacheData
	output string
}

var deserializeTestCases = []deserializeTestCase{
	// Existing test cases
	{
		input:  "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$12\r\nvalue\r\nvalue\r\n",
		output: []string{"SET", "key", "value\r\nvalue"},
	},
	{
		input:  "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
		output: []string{"SET", "key", "value"},
	},
	{
		input:  "*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n",
		output: []string{"GET", "key"},
	},
	{
		input:  "*3\r\n$6\r\nEXPIRE\r\n$3\r\nkey\r\n$2\r\n10\r\n",
		output: []string{"EXPIRE", "key", "10"},
	},
	{
		input:  "*2\r\n$3\r\nDEL\r\n$3\r\nkey\r\n",
		output: []string{"DEL", "key"},
	},
	{
		input:  "*2\r\n$4\r\nINCR\r\n$3\r\nkey\r\n",
		output: []string{"INCR", "key"},
	},
	{
		input:  "*3\r\n$5\r\nLPUSH\r\n$4\r\nlist\r\n$5\r\nvalue\r\n",
		output: []string{"LPUSH", "list", "value"},
	},
	{
		input:  "*4\r\n$6\r\nLRANGE\r\n$4\r\nlist\r\n$1\r\n0\r\n$2\r\n10\r\n",
		output: []string{"LRANGE", "list", "0", "10"},
	},
	{
		input:  "*2\r\n$4\r\nAUTH\r\n$5\r\nmyPwd\r\n",
		output: []string{"AUTH", "myPwd"},
	},
	{
		input:  "*1\r\n$4\r\nPING\r\n",
		output: []string{"PING"},
	},
	{
		input:  "*6\r\n$5\r\nHMSET\r\n$4\r\nhash\r\n$6\r\nfield1\r\n$6\r\nvalue1\r\n$6\r\nfield2\r\n$6\r\nvalue2\r\n",
		output: []string{"HMSET", "hash", "field1", "value1", "field2", "value2"},
	},
	{
		input:  "*0\r\n",
		output: []string{},
	},
	{
		input:  "*2\r\n$3\r\nSET\r\n$0\r\n\r\n",
		output: []string{"SET", ""},
	},

	// Additional commands:
	// MSET with multiple key-value pairs
	{
		input:  "*5\r\n$4\r\nMSET\r\n$4\r\nkey1\r\n$4\r\nval1\r\n$4\r\nkey2\r\n$4\r\nval2\r\n",
		output: []string{"MSET", "key1", "val1", "key2", "val2"},
	},
	// HSET for hash sets
	{
		input:  "*4\r\n$4\r\nHSET\r\n$4\r\nhash\r\n$5\r\nfield\r\n$5\r\nvalue\r\n",
		output: []string{"HSET", "hash", "field", "value"},
	},
	// ZADD for sorted sets
	{
		input:  "*4\r\n$4\r\nZADD\r\n$6\r\nmyzset\r\n$2\r\n10\r\n$5\r\nvalue\r\n",
		output: []string{"ZADD", "myzset", "10", "value"},
	},
	// Empty bulk string within a valid command
	{
		input:  "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$0\r\n\r\n",
		output: []string{"SET", "key", ""},
	},
	// Multi-key GET command
	{
		input:  "*3\r\n$4\r\nMGET\r\n$4\r\nkey1\r\n$4\r\nkey2\r\n",
		output: []string{"MGET", "key1", "key2"},
	},
	// WATCH for transactions
	{
		input:  "*2\r\n$5\r\nWATCH\r\n$3\r\nkey\r\n",
		output: []string{"WATCH", "key"},
	},
	// INCRBY with value
	{
		input:  "*3\r\n$6\r\nINCRBY\r\n$3\r\nkey\r\n$2\r\n10\r\n",
		output: []string{"INCRBY", "key", "10"},
	},
	// RPUSH for lists
	{
		input:  "*3\r\n$5\r\nRPUSH\r\n$4\r\nlist\r\n$5\r\nvalue\r\n",
		output: []string{"RPUSH", "list", "value"},
	},
	// EXPIRE with TTL
	{
		input:  "*3\r\n$6\r\nEXPIRE\r\n$3\r\nkey\r\n$2\r\n60\r\n",
		output: []string{"EXPIRE", "key", "60"},
	},
	// DEL for multiple keys
	{
		input:  "*3\r\n$3\r\nDEL\r\n$4\r\nkey1\r\n$4\r\nkey2\r\n",
		output: []string{"DEL", "key1", "key2"},
	},
	// SETNX to set only if key doesn't exist
	{
		input:  "*3\r\n$5\r\nSETNX\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
		output: []string{"SETNX", "key", "value"},
	},
	// SETEX with expiration
	{
		input:  "*4\r\n$5\r\nSETEX\r\n$3\r\nkey\r\n$2\r\n10\r\n$5\r\nvalue\r\n",
		output: []string{"SETEX", "key", "10", "value"},
	},
	// Empty array
	{
		input:  "*0\r\n",
		output: []string{},
	},
	// Invalid command
	{
		input:  "*2\r\n$7\r\nINVALID\r\n$3\r\nkey\r\n",
		output: []string{"INVALID", "key"},
	},
	{
		input:  "*3\r\n$3\r\nSET\r\n$4\r\nname\r\n$10\r\nShomi Khan\r\n",
		output: []string{"SET", "name", "Shomi Khan"},
	},
}

var serializeTestCases = []serializeTestCase{
	{
		input:  "SET key value",
		output: "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
	},
	{
		input:  "GET key",
		output: "*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n",
	},
	{
		input:  "PING",
		output: "*1\r\n$4\r\nPING\r\n",
	},
	{
		input:  "EXISTS key",
		output: "*2\r\n$6\r\nEXISTS\r\n$3\r\nkey\r\n",
	},
	{
		input:  "INCR counter",
		output: "*2\r\n$4\r\nINCR\r\n$7\r\ncounter\r\n",
	},
	{
		input:  "DECR counter",
		output: "*2\r\n$4\r\nDECR\r\n$7\r\ncounter\r\n",
	},
	{
		input:  "DEL key",
		output: "*2\r\n$3\r\nDEL\r\n$3\r\nkey\r\n",
	},
	{
		input:  "LPUSH mylist value1 value2",
		output: "*4\r\n$5\r\nLPUSH\r\n$6\r\nmylist\r\n$6\r\nvalue1\r\n$6\r\nvalue2\r\n",
	},
	{
		input:  "LPOP mylist",
		output: "*2\r\n$4\r\nLPOP\r\n$6\r\nmylist\r\n",
	},
	{
		input:  "RPUSH mylist value1 value2",
		output: "*4\r\n$5\r\nRPUSH\r\n$6\r\nmylist\r\n$6\r\nvalue1\r\n$6\r\nvalue2\r\n",
	},
	{
		input:  "RPOP mylist",
		output: "*2\r\n$4\r\nRPOP\r\n$6\r\nmylist\r\n",
	},
	{
		input:  "LRANGE mylist 0 -1",
		output: "*4\r\n$6\r\nLRANGE\r\n$6\r\nmylist\r\n$1\r\n0\r\n$2\r\n-1\r\n",
	},
	{
		input:  "EXPIRE key 60",
		output: "*3\r\n$6\r\nEXPIRE\r\n$3\r\nkey\r\n$2\r\n60\r\n",
	},
	{
		input:  "TTL key",
		output: "*2\r\n$3\r\nTTL\r\n$3\r\nkey\r\n",
	},
	{
		input:  "PERSIST key",
		output: "*2\r\n$7\r\nPERSIST\r\n$3\r\nkey\r\n",
	},
	{
		input:  "FLUSHALL",
		output: "*1\r\n$8\r\nFLUSHALL\r\n",
	},
}

var boolTestCases = []serializeTestCase{
	{
		input:  "true",
		output: ":1\r\n",
	},
	{
		input:  "false",
		output: ":0\r\n",
	},
}

var cacheItemTestCases = []serializeCacheDataTestCase{
	{
		input:  cache.CacheData{DataType: cache.Int, IntData: shared.IntToPtr(42)},
		output: ":42\r\n",
	},
	{
		input:  cache.CacheData{DataType: cache.String, StrData: shared.StringToPtr("hello")},
		output: "$5\r\nhello\r\n",
	},
	{
		input:  cache.CacheData{DataType: cache.Array, StrList: []string{"foo", "bar"}},
		output: "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
	},
	{
		input:  cache.CacheData{DataType: cache.Array, StrList: []string{}}, // Empty array
		output: "*0\r\n",
	},
	{
		input:  cache.CacheData{DataType: cache.String, StrData: nil}, // Nil string data
		output: "$-1\r\n",
	},
	{
		input:  cache.CacheData{DataType: cache.Int, IntData: nil}, // Nil int data
		output: "$-1\r\n",
	},
	{
		input:  cache.CacheData{DataType: cache.Array, StrList: nil}, // Nil array
		output: "$-1\r\n",
	},
}
