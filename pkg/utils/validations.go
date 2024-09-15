package utils

import (
	"strings"

	"github.com/nahK994/TinyCache/pkg/errors"
)

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

func ValidateSerializedCmd(serializedCmd string) error {
	errTypes := errors.GetErrorTypes()
	var cmdSegments []string
	numCmdSegments := 0
	if len(serializedCmd) == 0 || serializedCmd[0] != '*' {
		return errors.Err{Type: errTypes.UnexpectedCharacter}
	}
	index := 1 // Skip the initial '*'
	for {
		if index >= len(serializedCmd) {
			return errors.Err{Type: errTypes.IncompleteCommand}
		}

		ch := serializedCmd[index]
		if ch == '\r' {
			if !checkCRLF(serializedCmd, index) {
				return errors.Err{Type: errTypes.MissingCRLF}
			}
			index += 2 // Move past '\r\n'
			break
		}
		if !(ch >= '0' && ch <= '9') {
			return errors.Err{Type: errTypes.UnexpectedCharacter}
		}
		numCmdSegments = 10*numCmdSegments + int(ch-48)
		index++
	}

	size := 0
	for numCmdSegments != 0 {
		size = 0
		if serializedCmd[index] != '$' {
			return errors.Err{Type: errTypes.UnexpectedCharacter}
		}
		index++
		for {
			if index >= len(serializedCmd) {
				return errors.Err{Type: errTypes.IncompleteCommand}
			}

			ch := serializedCmd[index]
			if ch == '\r' {
				if !checkCRLF(serializedCmd, index) {
					return errors.Err{Type: errTypes.MissingCRLF}
				}
				index += 2
				break
			}
			if !(ch >= '0' && ch <= '9') {
				return errors.Err{Type: errTypes.InvalidBulkStringFormat}
			}

			size = 10*size + int(ch-48)
			index++
		}
		cmdSegments = append(cmdSegments, serializedCmd[index:index+size])

		if !checkCRLF(serializedCmd, index+size) {
			return errors.Err{Type: errTypes.MissingCRLF}
		}
		index += size + 2
		numCmdSegments--
	}

	if index != len(serializedCmd) {
		return errors.Err{Type: errTypes.SyntaxError}
	}
	if numCmdSegments != 0 {
		return errors.Err{Type: errTypes.IncompleteCommand}
	}

	return validateCmdArgs(cmdSegments)
}
