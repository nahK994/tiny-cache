package shared

import "strings"

func SplitCmd(rawCmd string) []string {
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

func PtrToInt(i int) *int {
	return &i
}

func PtrToString(s string) *string {
	return &s
}
