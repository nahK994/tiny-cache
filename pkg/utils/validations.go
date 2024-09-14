package utils

import (
	"github.com/nahK994/TinyCache/pkg/errors"
)

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

// TODO: Start work from here
// func ValidateRawCommand(rawCmd string) error {
// 	respCommands := GetRESPCommands()
// 	errType := errors.GetErrorTypes()
// 	words := GetCmdSegments(rawCmd)

// 	if len(words) == 0 {
// 		return errors.Err{Type: errType.UnknownCommand}
// 	}

// 	switch words[0] {
// 	case respCommands.SET:
// 		if len(words) < 3 {
// 			return errors.Err{Type: errType.WrongNumberOfArguments}
// 		}
// 	case respCommands.GET:
// 		if len(words) != 2 {
// 			return errors.Err{Type: errType.WrongNumberOfArguments}
// 		}
// 	case respCommands.EXISTS:
// 		if len(words) != 2 {
// 			return errors.Err{Type: errType.WrongNumberOfArguments}
// 		}
// 	case respCommands.DEL:
// 		if len(words) != 2 {
// 			return errors.Err{Type: errType.WrongNumberOfArguments}
// 		}
// 	case respCommands.INCR:
// 		if len(words) != 2 {
// 			return errors.Err{Type: errType.WrongNumberOfArguments}
// 		}
// 	case respCommands.DECR:
// 		if len(words) != 2 {
// 			return errors.Err{Type: errType.WrongNumberOfArguments}
// 		}
// 	default:
// 		return errors.Err{Type: errType.UnknownCommand}
// 	}

// 	return nil
// }

// func ValidateSerializedCmd(serializedCmd string) error {
// 	errTypes := errors.GetErrorTypes()

// 	numCmdSegments := 0
// 	index := 1 // Skip the initial '*'
// 	for {
// 		if index >= len(serializedCmd) {
// 			return errors.Err{Type: errTypes.IncompleteCommand}
// 		}

// 		ch := serializedCmd[index]
// 		if ch == '\r' {
// 			index += 2 // Move past '\r\n'
// 			break
// 		}
// 		if !(ch >= '0' && ch <= '9') {
// 			return errors.Err{Type: errTypes.UnexpectedCharacter}
// 		}
// 		numCmdSegments = 10*numCmdSegments + int(ch-48)
// 		index++
// 	}

// 	size := 0
// 	nextIdx := 0
// 	for i := 0; i < numCmdSegments; i++ {
// 		size = 0
// 		index++
// 		for {
// 			if index >= len(serializedCmd) {
// 				return errors.Err{Type: errTypes.IncompleteCommand}
// 			}

// 			ch := serializedCmd[index]
// 			if ch == '\r' {
// 				index += 2
// 				break
// 			}
// 			if !(ch >= '0' && ch <= '9') {
// 				fmt.Printf("%d %c %d\n", ch, rune(ch), index)
// 				return errors.Err{Type: errTypes.InvalidBulkStringFormat}
// 			}

// 			size = 10*size + int(ch-48)
// 			index++
// 		}

// 		if index+size > len(serializedCmd) || index+size+1 > len(serializedCmd) {
// 			return errors.Err{Type: errTypes.CommandLengthMismatch}
// 		}

// 		nextIdx = index + size
// 		if !(serializedCmd[nextIdx] == '\r' && serializedCmd[nextIdx+1] == '\n') {
// 			return errors.Err{Type: errTypes.MissingCRLF}
// 		}
// 		index = nextIdx + 2
// 	}

// 	return nil
// }
