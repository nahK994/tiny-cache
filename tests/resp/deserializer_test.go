package resp

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/resp"
)

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
