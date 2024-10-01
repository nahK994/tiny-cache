package validators

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/errors"
	"github.com/nahK994/TinyCache/pkg/validators"
)

var testCases = []struct {
	input       string
	expectedErr *errors.Err // Expecting either an error or nil
}{
	// Valid cases
	{
		input:       "SET key value\r\n",
		expectedErr: nil,
	},
	{
		input:       "GET key\r\n",
		expectedErr: nil,
	},
	{
		input:       "EXPIRE key 10",
		expectedErr: nil,
	},
	{
		input:       "EXPIRE key -10",
		expectedErr: &errors.Err{Type: errors.InvalidCommand},
	},
	{
		input:       "LRANGE mylist 0 10\r\n",
		expectedErr: nil,
	},
	{
		input:       "PING\r\n",
		expectedErr: nil,
	},
	// Valid case with multiple quotes, ensuring no invalid format error
	{
		input:       "SET \"quoted key\" \"quoted value\"\r\n",
		expectedErr: nil,
	},
	// Invalid command format (uneven number of quotes)
	{
		input:       "SET \"key value\r\n",
		expectedErr: &errors.Err{Type: errors.InvalidCommand},
	},
	// Unknown command
	{
		input:       "UNKNOWN key\r\n",
		expectedErr: &errors.Err{Type: errors.UnknownCommand},
	},
	// Incomplete command (missing arguments)
	{
		input:       "SET key\r\n",
		expectedErr: &errors.Err{Type: errors.IncompleteCommand},
	},
	// Too many arguments (for commands with max limit)
	{
		input:       "GET key extra_argument\r\n",
		expectedErr: &errors.Err{Type: errors.WrongNumberOfArguments},
	},
	// Wrong number of arguments (PING shouldn't have arguments)
	{
		input:       "PING extra\r\n",
		expectedErr: &errors.Err{Type: errors.WrongNumberOfArguments},
	},
	// Type error (expecting numeric value for EXPIRE)
	{
		input:       "EXPIRE key notANumber\r\n",
		expectedErr: &errors.Err{Type: errors.TypeError},
	},
	// Invalid command format with special characters
	{
		input:       "SET \"key\" \"val\r\n",
		expectedErr: &errors.Err{Type: errors.InvalidCommand},
	},
	// Invalid key with spaces
	{
		input:       "SET invalid key value\r\n",
		expectedErr: &errors.Err{Type: errors.TypeError},
	},
	// Correct command format but no key (edge case)
	{
		input:       "SET  \"\"\r\n",
		expectedErr: &errors.Err{Type: errors.IncompleteCommand},
	},
}

func TestValidateRawCommand(t *testing.T) {
	for _, tc := range testCases {
		err := validators.ValidateRawCommand(tc.input)

		if tc.expectedErr == nil {
			if err != nil {
				t.Errorf("Expected no error for input %q, but got %v", tc.input, err)
			}
		} else {
			if err == nil {
				t.Errorf("Expected error %v for input %q, but got no error", tc.expectedErr, tc.input)
			} else if typedErr, ok := err.(errors.Err); !ok || typedErr.Type != tc.expectedErr.Type {
				t.Errorf("Expected error type %q for input %q, but got %v", tc.expectedErr.Type, tc.input, err)
			}
		}
	}
}
