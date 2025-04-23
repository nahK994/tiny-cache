package resp

import (
	"errors"
)

func parseNumber(cmd string, index *int) int {
	numSegments := 0
	multiplier := 1
	if cmd[*index] == '-' {
		multiplier = -1
		*index++
	}
	for {
		ch := cmd[*index]
		if ch == '\r' {
			*index += 2
			break
		}

		numSegments = 10*numSegments + int(ch-48)
		*index++
	}
	return multiplier * numSegments
}

func Deserializer(rawCmd string) interface{} {
	index := 1
	typ := rune(rawCmd[0])

	switch typ {
	case '*':
		numSegments := parseNumber(rawCmd, &index)
		segments := make([]string, numSegments)
		for i := 0; i < numSegments; i++ {
			index++
			size := parseNumber(rawCmd, &index)
			segments[i] = rawCmd[index : index+size]
			index = index + size + 2
		}
		return segments
	case '$':
		var segment string
		length := parseNumber(rawCmd, &index)
		segment = rawCmd[index : index+length]
		return segment
	case ':':
		value := parseNumber(rawCmd, &index)
		return value
	case '+':
		return rawCmd[1 : len(rawCmd)-2]
	case '-':
		return errors.New(rawCmd[1:])
	default:
		return nil
	}
}
