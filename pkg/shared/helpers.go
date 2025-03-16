package shared

import (
	"strings"
)

func SplitCmd(rawCmd string) []string {
	list := strings.Fields(rawCmd)
	cmd := strings.Join(list, " ")
	isQuoteFound := false

	startIdx := -1
	endIdx := -1

	var segs []string

	for i := 0; i < len(cmd); i++ {
		if cmd[i] == '"' {
			if isQuoteFound {
				if startIdx == -1 && endIdx == -1 {
					continue
				}
				segs = append(segs, cmd[startIdx:endIdx+1])
				startIdx = -1
				endIdx = -1
			}

			isQuoteFound = !isQuoteFound
			continue
		}

		if !isQuoteFound {
			if cmd[i] != ' ' {
				if startIdx == -1 {
					startIdx = i
					endIdx = i
				} else {
					endIdx = i
				}
			} else {
				if startIdx == -1 {
					continue
				}
				segs = append(segs, cmd[startIdx:endIdx+1])
				startIdx = -1
				endIdx = -1
			}
		} else {
			if startIdx == -1 {
				startIdx = i
				endIdx = i
			} else {
				endIdx = i
			}
		}
	}

	if startIdx != -1 && endIdx != -1 {
		segs = append(segs, cmd[startIdx:endIdx+1])
	}

	return segs
}
