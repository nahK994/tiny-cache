package utils

import "github.com/nahK994/TinyCache/pkg/errors"

func ValidateRawCommand(rawCmd string) error {
	respCommands := GetRESPCommands()
	words := GetCmdSegments(rawCmd)
	switch words[0] {
	case respCommands.SET:
		if len(words) < 3 {
			return errors.InvalidCmd{Cmd: rawCmd}
		}
	case respCommands.GET:
		if len(words) != 2 {
			return errors.InvalidCmd{Cmd: rawCmd}
		}
	case respCommands.EXISTS:
		if len(words) != 2 {
			return errors.InvalidCmd{Cmd: rawCmd}
		}
	case respCommands.DEL:
		if len(words) != 2 {
			return errors.InvalidCmd{Cmd: rawCmd}
		}
	case respCommands.INCR:
		if len(words) != 2 {
			return errors.InvalidCmd{Cmd: rawCmd}
		}
	case respCommands.DECR:
		if len(words) != 2 {
			return errors.InvalidCmd{Cmd: rawCmd}
		}
	}
	return errors.InvalidCmd{Cmd: rawCmd}
}
