package utils

import (
	"strings"

	"github.com/nahK994/TinyCache/pkg/errors"
)

var errType errors.ErrTypes = errors.GetErrorTypes()

func validateCmdArgs(words []string) error {
	errType := errors.GetErrorTypes()
	switch strings.ToUpper(words[0]) {
	case respCommands.SET:
		if len(words) < 3 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCommands.GET:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCommands.EXISTS:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCommands.DEL:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCommands.INCR:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCommands.DECR:
		if len(words) != 2 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
		}
	case respCommands.PING:
		if len(words) != 1 {
			return errors.Err{Type: errType.WrongNumberOfArguments}
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
	*index++

	for {
		if *index >= len(cmd) {
			return -1, errors.Err{Type: errType.IncompleteCommand}
		}

		ch := cmd[*index]
		if ch == '\r' {
			if !checkCRLF(cmd, *index) {
				return -1, errors.Err{Type: errType.MissingCRLF}
			}
			*index += 2 // Move past '\r\n'
			break
		}
		if !(ch >= '0' && ch <= '9') {
			return -1, errors.Err{Type: errType.UnexpectedCharacter}
		}
		numCmdSegments = 10*numCmdSegments + int(ch-48)
		*index++
	}
	return numCmdSegments, nil
}

func getSegment(cmd string, index *int) (string, error) {
	if cmd[*index] != '$' {
		return "", errors.Err{Type: errType.UnexpectedCharacter}
	}
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
	index := 0
	var cmdSegments []string
	if len(serializedCmd) == 0 || serializedCmd[index] != '*' {
		return errors.Err{Type: errType.UnexpectedCharacter}
	}

	numCmdSegments, err := parseNumber(serializedCmd, &index)
	if err != nil {
		return err
	}

	for numCmdSegments != 0 {
		s, err1 := getSegment(serializedCmd, &index)
		if err1 != nil {
			return err1
		}
		cmdSegments = append(cmdSegments, s)
		numCmdSegments--
	}

	if index != len(serializedCmd) {
		return errors.Err{Type: errType.SyntaxError}
	}
	if numCmdSegments != 0 {
		return errors.Err{Type: errType.IncompleteCommand}
	}

	return validateCmdArgs(cmdSegments)
}
