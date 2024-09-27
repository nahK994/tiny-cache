package resp

import (
	"strings"

	"github.com/nahK994/TinyCache/pkg/errors"
)

func validateParseNumber(cmd string, index *int) (int, error) {
	numCmdSegments := 0
	for ; *index < len(cmd); *index++ {
		ch := cmd[*index]
		if *index+1 < len(cmd) && ch == '\r' && cmd[*index+1] == '\n' {
			*index += 2 // Move past '\r\n'
			return numCmdSegments, nil
		}

		if !(ch >= '0' && ch <= '9') {
			return -1, errors.Err{Type: errors.UnexpectedCharacter}
		}
		numCmdSegments = 10*numCmdSegments + int(ch-48)
	}
	return -1, errors.Err{Type: errors.IncompleteCommand}
}

func getSegment(cmd string, index *int) (string, error) {
	if cmd[*index] != '$' {
		return "", errors.Err{Type: errors.UnexpectedCharacter}
	}
	*index++
	size, err := validateParseNumber(cmd, index)
	if err != nil {
		return "", err
	}
	if !checkCRLF(cmd, *index+size) {
		return "", errors.Err{Type: errors.MissingCRLF}
	}
	seg := cmd[*index : *index+size]
	*index += (size + 2)
	return seg, nil
}

func checkCRLF(serializedCmd string, index int) bool {
	return index+1 < len(serializedCmd) && serializedCmd[index] == '\r' && serializedCmd[index+1] == '\n'
}

func getCmdSegments(rawCmd string) []string {
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
