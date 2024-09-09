package resp

import (
	"testing"
)

// Test basic deserialization of a simple RESP command
func TestDeserializer_Basic(t *testing.T) {
	respCmd := "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"
	expected := []string{"SET", "key", "value"}

	segments, err := Deserializer(respCmd)
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

// Test deserialization with different segment lengths
func TestDeserializer_VariedLength(t *testing.T) {
	respCmd := "*2\r\n$4\r\nPING\r\n$1\r\n1\r\n"
	expected := []string{"PING", "1"}

	segments, err := Deserializer(respCmd)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	for i, seg := range segments {
		if seg != expected[i] {
			t.Errorf("expected segment %d to be %s, got %s", i, expected[i], seg)
		}
	}
}

// Test deserialization with malformed input (missing end line)
func TestDeserializer_Malformed(t *testing.T) {
	respCmd := "*2\r\n$4\r\nPING\r\n$1\r\n1" // Missing the final \r\n

	_, err := Deserializer(respCmd)
	if err == nil {
		t.Fatalf("expected error for malformed input, got nil")
	}
}

// Test deserialization of a command with no segments
func TestDeserializer_EmptyCommand(t *testing.T) {
	respCmd := "*0\r\n"
	expected := []string{}

	segments, err := Deserializer(respCmd)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(segments) != len(expected) {
		t.Fatalf("expected %d segments, got %d", len(expected), len(segments))
	}
}
