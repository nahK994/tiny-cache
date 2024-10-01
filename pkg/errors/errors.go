package errors

type Err struct {
	Type string
}

func (e Err) Error() string {
	switch e.Type {
	case InvalidArrayFormat:
		return "-ERR Protocol error: invalid array format. Expected '*<count>\\r\\n'."

	case InvalidBulkStringFormat:
		return "-ERR Protocol error: invalid bulk string format. Expected '$<length>\\r\\n<string>\\r\\n'."

	case IncompleteCommand:
		return "-ERR Protocol error: incomplete command. Some parts of the serialized input are missing."

	case UnexpectedCharacter:
		return "-ERR Protocol error: unexpected character in command. Only valid characters are allowed."

	case CommandLengthMismatch:
		return "-ERR Protocol error: command length mismatch. Specified length does not match the actual data length."

	case UnknownCommand:
		return "-ERR unknown command"

	case WrongNumberOfArguments:
		return "-ERR wrong number of arguments for command"

	case SyntaxError:
		return "-ERR syntax error"

	case TypeError:
		return "-ERR Operation against a key holding the wrong kind of value"

	case CommandNotSupported:
		return "-ERR command not supported"

	case InvalidCommand:
		return "-ERR invalid command"

	case EmptyList:
		return "-ERR empty list or set"

	case UndefinedKey:
		return "-ERR key not defined"

	case ExpiredKey:
		return "-ERR key is expired"

	default:
		return "-ERR unknown command error"
	}
}
