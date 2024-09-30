package resp

import (
	"errors"
	"reflect"
	"testing"

	"github.com/nahK994/TinyCache/pkg/resp"
)

func TestDeserializer(t *testing.T) {
	for _, tc := range deserializeTestCases {
		result := resp.Deserializer(tc.input)

		if err, ok := result.(error); ok {
			expectedErr, isErr := tc.output.(error)
			if !isErr && !errors.Is(err, expectedErr) {
				t.Errorf("For input %q, expected error %v but got %v", tc.input, tc.output, result)
			}
		} else {
			if !reflect.DeepEqual(result, tc.output) {
				t.Errorf("For input %q, expected %v but got %v", tc.input, tc.output, result)
			}
		}
	}
}
