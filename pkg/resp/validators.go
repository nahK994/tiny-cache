package resp

import (
	"strconv"
	"strings"

	"github.com/nahK994/TinyCache/pkg/errors"
)

func validateCmdArgs(words []string) error {
	switch strings.ToUpper(words[0]) {
	case SET:
		if len(words) < 3 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case GET:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case EXISTS:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case DEL:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case INCR:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case DECR:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case PING:
		if len(words) != 1 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case FLUSHALL:
		if len(words) != 1 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case LPUSH:
		if len(words) < 3 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case LPOP:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case EXPIRE:
		if len(words) != 3 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case RPUSH:
		if len(words) < 3 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case RPOP:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case TTL:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case PERSIST:
		if len(words) != 2 {
			return errors.Err{Type: errors.WrongNumberOfArguments}
		}
	case LRANGE:
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
	words := getCmdSegments(rawCmd)

	if len(words) == 0 {
		return errors.Err{Type: errors.UnknownCommand}
	}

	return validateCmdArgs(words)
}

func ValidateSerializedCmd(serializedCmd string) error {
	if len(serializedCmd) == 0 {
		return errors.Err{Type: errors.IncompleteCommand}
	}

	index := 0
	if serializedCmd[index] != '*' {
		return errors.Err{Type: errors.UnexpectedCharacter}
	}
	index++
	numCmdSegments, err := validateParseNumber(serializedCmd, &index)
	if err != nil {
		return err
	}
	cmdSegments := make([]string, numCmdSegments)

	for index < len(serializedCmd) {
		if numCmdSegments == 0 {
			return errors.Err{Type: errors.CommandLengthMismatch}
		}

		s, err1 := getSegment(serializedCmd, &index)
		if err1 != nil {
			return err1
		}
		cmdSegments[len(cmdSegments)-numCmdSegments] = s
		numCmdSegments--
	}
	if numCmdSegments != 0 {
		return errors.Err{Type: errors.CommandLengthMismatch}
	}

	return validateCmdArgs(cmdSegments)
}
