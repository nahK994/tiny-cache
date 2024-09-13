package utils

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/utils"
)

func Test_MalformedRawCommands(t *testing.T) {
	for _, item := range malformedRawCmds {
		err := utils.ValidateRawCommand(item)
		if err == nil {
			t.Errorf("%s expected errors but no errors found", item)
		}
	}
}
