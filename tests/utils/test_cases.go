package utils

var malformedRawCmds []string = []string{
	"SET",
	"SET age ",
	"GET age val",
	"GET",
	"EXISTS",
	"EXISTS age val",
	"TEST",
	"PING haha",
	// Malformed INCR/DECR/DEL commands
	"INCR",         // Missing key argument
	"DECR",         // Missing key argument
	"DEL",          // Missing key argument
	"INCR age val", // Too many arguments for INCR
	"DECR age val", // Too many arguments for DECR
	"DEL key1 key2 key3",
}

var malformedSerializedCmds []string = []string{
	// Missing CRLF after array size
	"*2$3SET$3foo\r\n",
	// Explanation: The command declares an array with 2 segments, but it lacks the proper CRLF after the array size.

	// Incomplete bulk string (missing the string value)
	"*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$-1\r\n",
	// Explanation: Declares a bulk string, but it indicates a length of `-1` which is invalid.

	// Array count mismatch
	"*2\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
	// Explanation: The array count specifies 2, but the command provides 3 elements.

	// Unexpected character in array count
	"*x\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
	// Explanation: The array count has an invalid character `x`.

	// Missing CRLF after a bulk string
	"*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar",
	// Explanation: The final bulk string "bar" is missing a CRLF.

	// Command length mismatch
	"*2\r\n$3\r\nSET\r\n$5\r\nkeyvalue\r\n",
	// Explanation: The bulk string indicates a length of 5, but the actual string is longer.

	// Invalid array format (extra CRLF at the end)
	"*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n\r\n",
	// Explanation: There’s an extra CRLF at the end, which causes a mismatch in the length.

	// Missing bulk string length specifier
	"*3\r\n$3\r\nSET\r\nfoo\r\n$3\r\nbar\r\n",
	// Explanation: The length for the second bulk string "foo" is missing.

	// Invalid bulk string format (missing $ sign)
	"*3\r\n$3\r\nSET\r\n3\r\nfoo\r\n$3\r\nbar\r\n",
	// Explanation: The bulk string length for "foo" is missing the `$` sign.

	// Incomplete command (missing arguments)
	"*1\r\n$3\r\nSET\r\n",
	// Explanation: The array says 1 element, but `SET` requires at least 2 arguments.

	"-3\r\n$3\r\nSET\r\n$3\r\nage\r\n$3\r\n123\r\n",
	"*3\r\n-3\r\nSET\r\n$3\r\nage\r\n$3\r\n123\r\n",
	"*3\r\n$3\r\nSET\r\n-3\r\nage\r\n$3\r\n123\r\n",
	"*3\r\n$3\r\nSET\r\n$3\r\nage\r\n-3\r\n123\r\n",
	"*3\r\n$3\r\nSET\r\n$3\r\nage\r\n$3\r\a123\r\n",
	"*3\r\n$3\r\nGET\r\n$3\r\nage\r\n$3\r\n123\r\n",
}
