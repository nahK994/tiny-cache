package resp

func Deserializer(rawCmd string) []string {
	var segments []string
	numSegments := 0
	index := 1

	// Get number of command segments
	for {
		ch := rawCmd[index]
		if ch == '\r' {
			index += 2
			break
		}

		numSegments = 10*numSegments + int(ch-48)
		index++
	}

	size := 0
	var seg string
	for i := 0; i < numSegments; i++ {
		size = 0
		index++
		seg = ""
		for {
			ch := rawCmd[index]
			if ch == '\r' {
				index += 2
				break
			}

			size = 10*size + int(ch-48)
			index++
		}

		seg = rawCmd[index : index+size]
		segments = append(segments, seg)
		index = index + size + 2
	}
	return segments
}
