package utils

var malformedRawCmds []string = []string{
	"SET",
	"SET age ",
	"GET age val",
	"GET",
	"EXISTS",
	"EXISTS age val",
	"TEST",

	// Malformed INCR/DECR/DEL commands
	"INCR",         // Missing key argument
	"DECR",         // Missing key argument
	"DEL",          // Missing key argument
	"INCR age val", // Too many arguments for INCR
	"DECR age val", // Too many arguments for DECR
	"DEL key1 key2 key3",
}
