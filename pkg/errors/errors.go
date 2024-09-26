package errors

type Err struct {
	Type string
}

func (e Err) Error() string {
	errTypes := GetErrorTypes()

	switch e.Type {
	case errTypes.InvalidArrayFormat:
		return "-ERR Protocol error: invalid array format. Expected '*<count>\\r\\n'."

	case errTypes.InvalidBulkStringFormat:
		return "-ERR Protocol error: invalid bulk string format. Expected '$<length>\\r\\n<string>\\r\\n'."

	case errTypes.IncompleteCommand:
		return "-ERR Protocol error: incomplete command. Some parts of the serialized input are missing."

	case errTypes.UnexpectedCharacter:
		return "-ERR Protocol error: unexpected character in command. Only valid characters are allowed."

	case errTypes.MissingCRLF:
		return "-ERR Protocol error: missing CRLF (\\r\\n) after command segment."

	case errTypes.CommandLengthMismatch:
		return "-ERR Protocol error: command length mismatch. Specified length does not match the actual data length."

	case errTypes.UnknownCommand:
		return "-ERR unknown command"

	case errTypes.WrongNumberOfArguments:
		return "-ERR wrong number of arguments for command"

	case errTypes.SyntaxError:
		return "-ERR syntax error"

	case errTypes.TypeError:
		return "-ERR Operation against a key holding the wrong kind of value"

	case errTypes.CommandNotSupported:
		return "-ERR command not supported"

	case errTypes.InvalidCommandFormat:
		return "-ERR invalid command format"

	case errTypes.EmptyList:
		return "-ERR empty list or set"

	case errTypes.UndefinedKey:
		return "-ERR key not defined"

	case errTypes.ExpiredKey:
		return "-ERR key is expired"

	default:
		return "-ERR unknown command error"
	}
}
