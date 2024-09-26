package utils

import (
	"strconv"
	"strings"
	"time"

	"github.com/nahK994/TinyCache/pkg/errors"
)

func validateCmdArgs(words []string) error {
	errType := errors.GetErrorTypes()
	switch strings.ToUpper(words[0]) {
	case respCmds.SET:
		if len(words) < 3 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.GET:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.EXISTS:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.DEL:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.INCR:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.DECR:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.PING:
		if len(words) != 1 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.FLUSHALL:
		if len(words) != 1 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.LPUSH:
		if len(words) < 3 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.LPOP:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.EXPIRE:
		if len(words) != 3 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.RPUSH:
		if len(words) < 3 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.RPOP:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.TTL:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.PERSIST:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.LRANGE:
		if len(words) != 4 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
		_, strIdx_ok := strconv.Atoi(words[2])
		_, endIdx_ok := strconv.Atoi(words[3])
		if strIdx_ok != nil || endIdx_ok != nil {
			return errors.Err{Type: errType.TypeError}
		}
	default:
		return errors.Err{Type: errType.UnknownCommand}
	}
	return nil
}

func ValidateRawCommand(rawCmd string) error {
	errType := errors.GetErrorTypes()
	words := GetCmdSegments(rawCmd)

	if len(words) == 0 {
		return errors.Err{Type: errType.UnknownCommand}
	}

	return validateCmdArgs(words)
}

func ValidateSerializedCmd(serializedCmd string) error {
	if len(serializedCmd) == 0 {
		return errors.Err{Type: errType.IncompleteCommand}
	}

	index := 0
	if serializedCmd[index] != '*' {
		return errors.Err{Type: errType.UnexpectedCharacter}
	}
	index++
	numCmdSegments, err := parseNumber(serializedCmd, &index)
	if err != nil {
		return err
	}
	cmdSegments := make([]string, numCmdSegments)

	for index < len(serializedCmd) {
		if numCmdSegments == 0 {
			return errors.Err{Type: errType.CommandLengthMismatch}
		}

		s, err1 := getSegment(serializedCmd, &index)
		if err1 != nil {
			return err1
		}
		cmdSegments[len(cmdSegments)-numCmdSegments] = s
		numCmdSegments--
	}
	if numCmdSegments != 0 {
		return errors.Err{Type: errType.CommandLengthMismatch}
	}

	return validateCmdArgs(cmdSegments)
}

func CheckEmptyList(key string) error {
	if IsKeyExists(key) {
		return errors.Err{Type: errType.EmptyList}
	}
	return nil
}

func validateExpiry(key string) error {
	item := c.GET(key)
	if item.ExpiryTime != nil && time.Now().After(*item.ExpiryTime) {
		c.DEL(key)
		return errors.Err{Type: errType.ExpiredKey}
	}
	return nil
}

func AssertKeyExists(key string) error {
	if err := validateExpiry(key); err != nil {
		return err
	}
	if !c.EXISTS(key) {
		return errors.Err{Type: errType.UndefinedKey}
	}
	return nil
}

func AssertListType(key string) error {
	if _, ok := c.GET(key).Val.([]string); !ok {
		return errors.Err{Type: errType.TypeError}
	}
	return nil
}

func AssertIntType(key string) error {
	if _, ok := c.GET(key).Val.(int); !ok {
		return errors.Err{Type: errType.TypeError}
	}
	return nil
}
