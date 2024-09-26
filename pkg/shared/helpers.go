package shared

import (
	"strings"
)

func GetCmdSegments(rawCmd string) []string {
	var words []string
	temp := strings.Split(rawCmd, " ")
	for _, ch := range temp {
		if len(ch) == 0 {
			continue
		}
		words = append(words, ch)
	}
	return words
}
