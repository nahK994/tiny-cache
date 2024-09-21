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

		segments, _ := resp.Deserializer(respCmd).([]string)
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
