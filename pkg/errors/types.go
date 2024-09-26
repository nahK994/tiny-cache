package errors

type ErrTypes struct {
	InvalidArrayFormat      string
	InvalidBulkStringFormat string
	IncompleteCommand       string
	UnexpectedCharacter     string
	MissingCRLF             string
	CommandLengthMismatch   string
	UnknownCommand          string
	WrongNumberOfArguments  string
	SyntaxError             string
	TypeError               string
	CommandNotSupported     string
	InvalidCommandFormat    string
	EmptyList               string
	UndefinedKey            string
	ExpiredKey              string
}

var errType ErrTypes = ErrTypes{
	InvalidArrayFormat:      "InvalidArrayFormat",
	InvalidBulkStringFormat: "InvalidBulkStringFormat",
	IncompleteCommand:       "IncompleteCommand",
	UnexpectedCharacter:     "UnexpectedCharacter",
	MissingCRLF:             "MissingCRLF",
	CommandLengthMismatch:   "CommandLengthMismatch",
	UnknownCommand:          "UnknownCommand",
	WrongNumberOfArguments:  "WrongNumberOfArguments",
	SyntaxError:             "SyntaxError",
	TypeError:               "TypeError",
	CommandNotSupported:     "CommandNotSupported",
	InvalidCommandFormat:    "InvalidCommandFormat",
	EmptyList:               "EmptyList",
	UndefinedKey:            "UndefinedKey",
	ExpiredKey:              "ExpiredKey",
}

func GetErrorTypes() ErrTypes {
	return errType
}
