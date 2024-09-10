package resp

import (
	"strconv"
	"strings"
)

func getCmdSegments(cmd string) []string {
	temp := strings.Split(cmd, " ")
	var ans []string
	for _, ch := range temp {
		if len(ch) == 0 {
			continue
		}
		ans = append(ans, ch)
	}

	return ans
}

func Serialize(cmd string) (string, error) {
	if cmd == "" {
		return "*0\r\n", nil
	}
	if cmd == " " {
		return "*1\r\n$1\r\n \r\n", nil
	}

	segments := getCmdSegments(cmd)
	serializedCmd := "*" + strconv.Itoa(len(segments)) + "\r\n"

	for _, seg := range segments {
		serializedCmd += ("$" + strconv.Itoa(len(seg)) + "\r\n" + seg + "\r\n")
	}

	return serializedCmd, nil
}
