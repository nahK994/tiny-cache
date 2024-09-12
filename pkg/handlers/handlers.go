package handlers

import (
	"fmt"
	"strings"

	"github.com/nahK994/TinyCache/pkg/cache"
	"github.com/nahK994/TinyCache/pkg/errors"
	"github.com/nahK994/TinyCache/pkg/resp"
	"github.com/nahK994/TinyCache/pkg/utils"
)

var c *cache.Cache = cache.InitCache()

func handleGET(segments []string) (string, error) {
	replytype := utils.GetReplyTypes()
	if len(segments) > 1 {
		return "", errors.Err{Msg: "-ERR wrong number of arguments for 'GET' command\r\n", File: "handlers/handlers.go", Line: 19}
	}
	if !c.IsKeyExist(segments[0]) {
		return "$-1\r\n", nil
	}

	val, ok := c.ReadCache(segments[0]).(string)
	if !ok {
		return "", errors.Err{Msg: "-Unknown datatype\r\n", File: "handlers/handlers.go", Line: 26}
	}
	return fmt.Sprintf("%c%d\r\n%s\r\n", replytype.Bulk, len(val), val), nil
}

func handleSET(segments []string) (string, error) {
	key := segments[0]
	value := segments[1]
	if err := c.WriteCache(key, value); err != nil {
		return "", err
	}
	return "+OK\r\n", nil
}

func handleKeyExist(segments []string) (string, error) {
	if len(segments) > 1 || len(segments) < 1 {
		return "", errors.Err{Msg: "-ERR unknown command 'INVALID_COMMAND'\r\n", File: "handlers/handlers.go", Line: 42}
	}

	if c.IsKeyExist(segments[0]) {
		return ":1\r\n", nil
	} else {
		return ":0\r\n", nil
	}
}

func HandleCommand(serializedCmd string) (string, error) {
	cmdSegments, err := resp.Deserializer(serializedCmd)
	respCmd := utils.GetRESPCommands()
	if err != nil {
		return "", err
	}

	switch strings.ToUpper(cmdSegments[0]) {
	case respCmd.GET:
		return handleGET(cmdSegments[1:])
	case respCmd.SET:
		return handleSET(cmdSegments[1:])
	case respCmd.EXISTS:
		return handleKeyExist(cmdSegments[1:])
	case respCmd.PING:
		return "+PONG\r\n", nil
	default:
		return fmt.Sprintln("Please use these commands:", strings.Join([]string{
			respCmd.SET, respCmd.GET, respCmd.EXISTS, respCmd.PING,
		}, ", ")), nil
	}
}
