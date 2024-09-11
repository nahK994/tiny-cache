package handlers

import (
	"fmt"
	"strconv"
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
		return "", errors.MalformedErr{Msg: "-ERR wrong number of arguments for 'GET' command\r\n"}
	}
	if !c.IsKeyExist(segments[0]) {
		return "", errors.MalformedErr{Msg: "-Key not exists\r\n"}
	}

	val := c.ReadCache(segments[0])
	switch v := val.(type) {
	case int:
		return fmt.Sprint(replytype.Int, v, "\r\n"), nil
	case string:
		return fmt.Sprint(replytype.Bulk, v, "\r\n"), nil
	case []string:
		resp := fmt.Sprint(replytype.Array, strconv.Itoa(len(v)), "\r\n")
		for _, item := range v {
			resp += fmt.Sprint(replytype.Bulk, item, "\r\n")
		}
		return resp, nil
	default:
		return "", errors.MalformedErr{Msg: "-Unknown datatype\r\n"}
	}
}

func handleSET(segments []string) (string, error) {
	key := segments[0]
	value := strings.Join(segments[1:], " ")
	if err := c.WriteCache(key, value); err != nil {
		return "", err
	}
	return "+OK\r\n", nil
}

func handleKeyExist(segments []string) (string, error) {
	if len(segments) > 1 || len(segments) < 1 {
		return "", errors.MalformedErr{Msg: "-ERR unknown command 'INVALID_COMMAND'\r\n"}
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
	default:
		return fmt.Sprintln("Please use these commands:", strings.Join([]string{
			respCmd.SET, respCmd.GET,
		}, ", ")), nil
	}
}
