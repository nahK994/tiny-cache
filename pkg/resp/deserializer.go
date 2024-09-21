package resp

import "github.com/nahK994/TinyCache/pkg/utils"

func parseNumber(cmd string, index *int) int {
	numSegments := 0
	for {
		ch := cmd[*index]
		if ch == '\r' {
			*index += 2
			break
		}

		numSegments = 10*numSegments + int(ch-48)
		*index++
	}
	return numSegments
}

func Deserializer(rawCmd string) interface{} {
	index := 1
	types := utils.GetReplyTypes()
	typ := rune(rawCmd[0])

	switch typ {
	case types.Array:
		var segments []string
		numSegments := parseNumber(rawCmd, &index)
		for i := 0; i < numSegments; i++ {
			index++
			size := parseNumber(rawCmd, &index)
			segments = append(segments, rawCmd[index:index+size])
			index = index + size + 2
		}
		return segments
	case types.Bulk:
		var segment string
		length := parseNumber(rawCmd, &index)
		segment = rawCmd[index : index+length]
		return segment
	case types.Int:
		value := parseNumber(rawCmd, &index)
		return value
	case types.Status:
		return rawCmd[1 : len(rawCmd)-2]
	default:
		return nil
	}

}
