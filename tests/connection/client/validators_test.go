package client

import (
	"testing"

	"github.com/nahK994/TinyCache/connection/client"
)

var malformedRawCmds []string = []string{
	"SET",                        // Missing value for the SET command
	"SET key",                    // Missing value for the SET command
	"GET",                        // Missing key for the GET command
	"EXISTS",                     // Missing key for the EXISTS command
	"DEL",                        // Missing key for the DEL command
	"INCR",                       // Missing key for the INCR command
	"DECR",                       // Missing key for the DECR command
	"PING arg",                   // PING command should not have any arguments
	"FLUSHALL arg",               // FLUSHALL should not have any arguments
	"LPUSH key",                  // Missing value for LPUSH
	"LPOP",                       // Missing key for LPOP
	"EXPIRE key",                 // Missing TTL for EXPIRE
	"RPUSH key",                  // Missing value for RPUSH
	"RPOP",                       // Missing key for RPOP
	"TTL",                        // Missing key for TTL
	"PERSIST",                    // Missing key for PERSIST
	"LRANGE key start",           // Missing end index for LRANGE
	"LRANGE key start end extra", // Extra argument for LRANGE
	"LRANGE key nonInt 2",        // Non-integer start index for LRANGE
	"LRANGE key 2 nonInt",        // Non-integer end index for LRANGE
	"UNKNOWN",                    // Unknown command
	"EXPIRE key nonIntTTL",       // Non-integer TTL for EXPIRE
	"LPOP key extraArg",          // Extra argument for LPOP
}

func Test_MalformedRawCommands(t *testing.T) {
	for _, item := range malformedRawCmds {
		err := client.ValidateRawCommand(item)
		if err == nil {
			t.Errorf("%s expected errors but no errors found", item)
		}
	}
}
