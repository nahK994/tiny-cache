package resp

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/resp"
)

func TestSerialize(t *testing.T) {
	for _, item := range serializeTestCases {
		serialized := resp.Serialize(item.input)

		if serialized != item.output {
			t.Errorf("input = %s expected %s, got %s", item.input, item.output, serialized)
		}
	}
}
