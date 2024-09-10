package resp

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/resp"
)

type deserializeTestCase struct {
	input  string
	output []string
}

var deserializeTestCases = []deserializeTestCase{
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
}

// Test basic deserialization of a simple RESP command
func TestDeserializer_Basic(t *testing.T) {
	for _, item := range deserializeTestCases {
		respCmd := item.input
		expected := item.output

		segments, _ := resp.Deserializer(respCmd)
		if len(segments) != len(expected) {
			t.Fatalf("expected %d segments, got %d", len(expected), len(segments))
		}

		for i, seg := range segments {
			if seg != expected[i] {
				t.Errorf("expected segment at position %d to be %s, got %s", i, expected[i], seg)
			}
		}
	}
}

// Test deserialization with malformed input (missing end line)
func TestDeserializer_Malformed(t *testing.T) {
	respCmd := "*2\r\n$4\r\nPING\r\n$1\r\n1" // Missing the final \r\n

	_, err := resp.Deserializer(respCmd)
	if err == nil {
		t.Fatalf("expected error for malformed input, got nil")
	}
}

func TestDeserializer_EmptyInput(t *testing.T) {
	respCmd := ""

	_, err := resp.Deserializer(respCmd)
	if err == nil {
		t.Fatalf("expected error for empty input, got nil")
	}
}

func TestDeserializer_IncorrectArraySize(t *testing.T) {
	respCmd := "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n" // Missing third element

	_, err := resp.Deserializer(respCmd)
	if err == nil {
		t.Fatalf("expected error for incorrect array size, got nil")
	}
}

func TestDeserializer_NegativeBulkStringLength(t *testing.T) {
	respCmd := "*2\r\n$-3\r\nSET\r\n$3\r\nkey\r\n" // Invalid negative length for bulk string

	_, err := resp.Deserializer(respCmd)
	if err == nil {
		t.Fatalf("expected error for negative bulk string length, got nil")
	}
}

func TestDeserializer_EmptyBulkString(t *testing.T) {
	respCmd := "*2\r\n$3\r\nSET\r\n$0\r\n\r\n" // Empty bulk string

	expected := []string{"SET", ""}
	segments, err := resp.Deserializer(respCmd)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(segments) != len(expected) {
		t.Fatalf("expected %d segments, got %d", len(expected), len(segments))
	}

	for i, seg := range segments {
		if seg != expected[i] {
			t.Errorf("expected segment %d to be %s, got %s", i, expected[i], seg)
		}
	}
}
