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

func ValidateSerializedCmd(serializedCmd string) error {
	errTypes := errors.GetErrorTypes()
	var cmdSegments []string
	numCmdSegments := 0
	index := 1 // Skip the initial '*'
	for {
		if index >= len(serializedCmd) {
			return errors.Err{Type: errTypes.IncompleteCommand}
		}

		ch := serializedCmd[index]
		if ch == '\r' {
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
	var i int
	for i = 0; i < numCmdSegments; i++ {
		size = 0
		index++
		for {
			if index >= len(serializedCmd) {
				return errors.Err{Type: errTypes.IncompleteCommand}
			}

			ch := serializedCmd[index]
			if ch == '\r' {
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

		if index+size >= len(serializedCmd) || index+size+1 >= len(serializedCmd) {
			return errors.Err{Type: errTypes.CommandLengthMismatch}
		}

		if !(serializedCmd[index+size] == '\r' && serializedCmd[index+size+1] == '\n') {
			return errors.Err{Type: errTypes.MissingCRLF}
		}
		index += size + 2
	}

	if index != len(serializedCmd) {
		return errors.Err{Type: errTypes.SyntaxError}
	}
	if i != numCmdSegments {
		return errors.Err{Type: errTypes.IncompleteCommand}
	}

	return validateCmdArgs(cmdSegments)
}
