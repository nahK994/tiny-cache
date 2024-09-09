package resp

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/resp"
)

type serializeTestCase struct {
	input  string
	output string
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
		input:  "",
		output: "*0\r\n",
	},
	{
		input:  "PING",
		output: "*1\r\n$4\r\nPING\r\n",
	},
	{
		input:  "HMSET hash field1 value1 field2 value2",
		output: "*6\r\n$5\r\nHMSET\r\n$4\r\nhash\r\n$6\r\nfield1\r\n$6\r\nvalue1\r\n$6\r\nfield2\r\n$6\r\nvalue2\r\n",
	},
	{
		input:  "LRANGE list 0 10",
		output: "*4\r\n$6\r\nLRANGE\r\n$4\r\nlist\r\n$1\r\n0\r\n$2\r\n10\r\n",
	},
	{
		input:  "AUTH myPwd",
		output: "*2\r\n$4\r\nAUTH\r\n$5\r\nmyPwd\r\n",
	},
}

func TestSerialize(t *testing.T) {
	for _, item := range serializeTestCases {
		serialized, err := resp.Serialize(item.input)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if serialized != item.output {
			t.Errorf("expected %s, got %s", item.output, serialized)
		}
	}
}
