package resp

import "errors"

type segmentLength int
type nextIndex int
type segment string

func getSegmentLength(cmd string, startIndex int) (segmentLength, nextIndex) {
	length := 0
	startIndex++ // Move past the '$' or '*' character
	for {
		if cmd[startIndex] == '\r' {
			break
		}

		length = 10*length + int(cmd[startIndex]-48)
		startIndex++
	}

	return segmentLength(length), nextIndex(startIndex + 1)
}

func getSegment(cmd string, startIndex int, segLength int) (segment, nextIndex) {
	seg := ""
	startIndex++ // Move past the initial '$'
	for i := startIndex; i < startIndex+segLength; i++ {
		seg += string(cmd[i])
	}
	return segment(seg), nextIndex(startIndex + segLength + 2) // Skip '\r\n' after segment
}

func Deserializer(cmd string) ([]string, error) {
	var segments []string
	numSegments := 0
	index := 1 // Skip the initial '*'

	// Read number of segments
	for {
		if cmd[index] == '\r' {
			index += 2 // Move past '\r\n'
			break
		}
		numSegments = 10*numSegments + int(cmd[index]-48)
		index++
	}

	// Read each segment based on the given number of segments
	for numSegments > 0 {
		length, nextIdx := getSegmentLength(cmd, index)
		seg, nextIdx := getSegment(cmd, int(nextIdx), int(length))

		index = int(nextIdx)
		segments = append(segments, string(seg))
		numSegments--
	}

	if index != len(cmd) {
		return nil, errors.New("malformed error")
	}
	return segments, nil
}
