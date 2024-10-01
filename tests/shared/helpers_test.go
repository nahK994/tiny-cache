package shared

import (
	"reflect"
	"testing"

	"github.com/nahK994/TinyCache/pkg/shared"
)

type testCaseType struct {
	input  string
	output []string
}

var testCases = []testCaseType{
	{
		input:  "set name \"shomi khan\"",
		output: []string{"set", "name", "shomi khan"},
	},
	{
		input:  "lpush arr \"shomi khan\" \"bob ross\"",
		output: []string{"lpush", "arr", "shomi khan", "bob ross"},
	},
	{
		input:  "lrange arr 0 10",
		output: []string{"lrange", "arr", "0", "10"},
	},
}

func TestSplitCmd(t *testing.T) {
	for _, tc := range testCases {
		got := shared.SplitCmd(tc.input)
		if !reflect.DeepEqual(got, tc.output) {
			t.Errorf("input %s Expected = %v; got %v", tc.input, tc.output, got)
		}
	}
}
