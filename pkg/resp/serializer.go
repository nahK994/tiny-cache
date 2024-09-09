package resp

import (
	"strconv"
	"strings"
)

func Serialize(cmd string) (string, error) {
	if cmd == "" {
		return "*0\r\n", nil
	}

	segments := strings.Split(cmd, " ")
	serializedCmd := "*" + strconv.Itoa(len(segments)) + "\r\n"

	for _, seg := range segments {
		serializedCmd += ("$" + strconv.Itoa(len(seg)) + "\r\n" + seg + "\r\n")
	}

	return serializedCmd, nil
}
