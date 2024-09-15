package resp

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

func Deserializer(rawCmd string) []string {
	var segments []string
	index := 1
	numSegments := parseNumber(rawCmd, &index)

	for i := 0; i < numSegments; i++ {
		index++
		size := parseNumber(rawCmd, &index)
		segments = append(segments, rawCmd[index:index+size])
		index = index + size + 2
	}
	return segments
}
