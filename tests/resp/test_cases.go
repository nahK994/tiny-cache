package resp

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/nahK994/tiny-cache/pkg/cache"
	"github.com/nahK994/tiny-cache/pkg/utils"
)

type deserializeTestCase struct {
	input  string
	output interface{}
}

type serializeTestCase struct {
	input  string
	output string
}

type serializeCacheDataTestCase struct {
	input  cache.DataItem
	output string
}

type serializeBoolTestCase struct {
	input  bool
	output string
}

var deserializeTestCases = []deserializeTestCase{
	{
		input:  "*2\r\n$3\r\nGET\r\n$4\r\nname\r\n",
		output: []string{"GET", "name"},
	},
	{
		input:  "$3\r\nfoo\r\n",
		output: "foo",
	},
	{
		input:  ":1000\r\n",
		output: 1000,
	},
	{
		input:  "+OK\r\n",
		output: "OK",
	},
	{
		input:  "-ERR unknown command\r\n",
		output: errors.New("ERR unknown command"),
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

var boolTestCases = []serializeBoolTestCase{
	{
		input:  true,
		output: ":1\r\n",
	},
	{
		input:  false,
		output: ":0\r\n",
	},
}

func strListToByteSlice(list []string) []byte {
	ls, _ := json.Marshal(list)
	return ls
}

var cacheItemTestCases = []serializeCacheDataTestCase{
	{
		input:  cache.DataItem{DataType: utils.Int, Value: []byte(strconv.Itoa(42))},
		output: ":42\r\n",
	},
	{
		input:  cache.DataItem{DataType: utils.String, Value: []byte("hello")},
		output: "$5\r\nhello\r\n",
	},
	{
		input:  cache.DataItem{DataType: utils.Array, Value: strListToByteSlice([]string{"foo", "bar"})},
		output: "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
	},
	{
		input:  cache.DataItem{DataType: utils.Array, Value: strListToByteSlice([]string{})}, // Empty array
		output: "*0\r\n",
	},
	// {
	// 	input:  cache.DataItem{DataType: utils.String, Value: []byte(nil)}, // Nil string data
	// 	output: "$-1\r\n",
	// },
	{
		input:  cache.DataItem{DataType: utils.Int, Value: nil}, // Nil int data
		output: "$-1\r\n",
	},
	{
		input:  cache.DataItem{DataType: utils.Array, Value: []byte{}}, // Nil array
		output: "$-1\r\n",
	},
}
