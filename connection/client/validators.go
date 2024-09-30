package client

import (
	"strconv"
	"strings"

	"github.com/nahK994/TinyCache/pkg/errors"
	"github.com/nahK994/TinyCache/pkg/resp"
	"github.com/nahK994/TinyCache/pkg/shared"
)

func validateCmdArgs(words []string) error {
	switch strings.ToUpper(words[0]) {
	case resp.SET:
		if len(words) < 3 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.GET:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.EXISTS:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.DEL:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.INCR:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.DECR:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.PING:
		if len(words) != 1 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.FLUSHALL:
		if len(words) != 1 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.LPUSH:
		if len(words) < 3 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.LPOP:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.EXPIRE:
		if len(words) != 3 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
		if _, err := strconv.Atoi(words[2]); err != nil {
			return errors.Err{Type: errors.TypeError}
		}
	case resp.RPUSH:
		if len(words) < 3 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.RPOP:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.TTL:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.PERSIST:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case resp.LRANGE:
		if len(words) != 4 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
		_, strIdx_ok := strconv.Atoi(words[2])
		_, endIdx_ok := strconv.Atoi(words[3])
		if strIdx_ok != nil || endIdx_ok != nil {
			return errors.Err{Type: errors.TypeError}
		}
	default:
		return errors.Err{Type: errors.UnknownCommand}
	}
	return nil
}

func ValidateRawCommand(rawCmd string) error {
	words := shared.SplitCmd(rawCmd)

	if len(words) == 0 {
		return errors.Err{Type: errors.UnknownCommand}
	}

	return validateCmdArgs(words)
}
