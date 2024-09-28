package client

import (
	"testing"

	"github.com/nahK994/TinyCache/connection/client"
)

func Test_MalformedRawCommands(t *testing.T) {
	for _, item := range malformedRawCmds {
		err := client.ValidateRawCommand(item)
		if err == nil {
			t.Errorf("%s expected errors but no errors found", item)
		}
	}
}
