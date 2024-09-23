package utils

import (
	"strconv"
	"strings"

	"github.com/nahK994/TinyCache/pkg/errors"
)

var errType errors.ErrTypes = errors.GetErrorTypes()

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
	case respCmds.RPUSH:
		if len(words) < 3 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCmds.RPOP:
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

func checkCRLF(serializedCmd string, index int) bool {
	return index+1 < len(serializedCmd) && serializedCmd[index] == '\r' && serializedCmd[index+1] == '\n'
}

func parseNumber(cmd string, index *int) (int, error) {
	numCmdSegments := 0
	for ; *index < len(cmd); *index++ {
		ch := cmd[*index]
		if *index+1 < len(cmd) && ch == '\r' && cmd[*index+1] == '\n' {
			*index += 2 // Move past '\r\n'
			return numCmdSegments, nil
		}

		if !(ch >= '0' && ch <= '9') {
			return -1, errors.Err{Type: errType.UnexpectedCharacter}
		}
		numCmdSegments = 10*numCmdSegments + int(ch-48)
	}
	return -1, errors.Err{Type: errType.IncompleteCommand}
}

func getSegment(cmd string, index *int) (string, error) {
	if cmd[*index] != '$' {
		return "", errors.Err{Type: errType.UnexpectedCharacter}
	}
	*index++
	size, err := parseNumber(cmd, index)
	if err != nil {
		return "", err
	}
	if !checkCRLF(cmd, *index+size) {
		return "", errors.Err{Type: errType.MissingCRLF}
	}
	seg := cmd[*index : *index+size]
	*index += (size + 2)
	return seg, nil
}

func ValidateSerializedCmd(serializedCmd string) error {
	var cmdSegments []string
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

	for index < len(serializedCmd) {
		if len(cmdSegments) == numCmdSegments {
			return errors.Err{Type: errType.CommandLengthMismatch}
		}

		s, err1 := getSegment(serializedCmd, &index)
		if err1 != nil {
			return err1
		}
		cmdSegments = append(cmdSegments, s)
	}

	err2 := validateCmdArgs(cmdSegments)
	return err2
}
