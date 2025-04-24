package resp

import (
	"testing"

	"github.com/nahK994/tiny-cache/pkg/resp"
)

func TestSerialize(t *testing.T) {
	for _, item := range serializeTestCases {
		serialized := resp.Serialize(item.input)
		if serialized != item.output {
			t.Errorf("input = %s expected %s, got %s", item.input, item.output, serialized)
		}
	}
}

func TestSerializeBool(t *testing.T) {
	for _, tc := range boolTestCases {
		got := resp.SerializeBool(tc.input)
		if got != tc.output {
			t.Errorf("SerializeBool(%v) = %v; want %v", tc.input, got, tc.output)
		}
	}
}

func TestSerializeCacheItem(t *testing.T) {
	for _, tc := range cacheItemTestCases {
		got := resp.SerializeCacheItem(&tc.input)
		if got != tc.output {
			t.Errorf("SerializeCacheItem(%v) = %v; want %v", tc.input, got, tc.output)
		}
	}
}
