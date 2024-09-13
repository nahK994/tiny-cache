package resp

import (
	"fmt"

	"github.com/nahK994/TinyCache/pkg/errors"
)

type segmentLength int
type nextIndex int
type segment string

func isSegEnded(cmd string, index int) (bool, error) {
	if index+1 >= len(cmd) {
		return false, errors.Err{Msg: fmt.Sprintln("cmd ended before segment ended"), File: "resp/deserializer.go", Line: 15}
	}
	return cmd[index] == '\r' && cmd[index+1] == '\n', nil
}

func getSegmentLength(cmd string, startIndex int) (segmentLength, nextIndex, error) {
	length := 0
	startIndex++ // Move past the '$' character
	for {
		isSegEnded, err := isSegEnded(cmd, startIndex)
		if err != nil {
			return -1, -1, err
		}
		if isSegEnded {
			return segmentLength(length), nextIndex(startIndex + 1), nil
		}

		ch := cmd[startIndex]
		if !(ch >= '0' && ch <= '9') {
			return -1, -1, errors.Err{Msg: fmt.Sprintf("Malformed error from getSegmentLength for %v %d", cmd[startIndex], startIndex)}
		}

		length = 10*length + int(ch-48)
		startIndex++
	}
}

func getSegment(cmd string, startIndex int, segLength int) (segment, nextIndex, error) {
	_, err := isSegEnded(cmd, startIndex+segLength+1)
	if err != nil {
		return "", -1, err
	}

	seg := ""
	startIndex++ // Move past the initial '$'
	for i := startIndex; i < startIndex+segLength; i++ {
		seg += string(cmd[i])
	}
	return segment(seg), nextIndex(startIndex + segLength + 2), nil // Skip '\r\n' after segment
}

func Deserializer(rawCmd string) ([]string, error) {
	var segments []string
	numSegments := 0
	index := 1 // Skip the initial '*'

	if len(rawCmd) < 2 {
		return nil, errors.Err{Msg: fmt.Sprintf("Malformed error from Deserializer for %s\n", rawCmd), File: "resp/deserializer.go", Line: 62}
	}

	// Read number of segments
	for {
		ch := rawCmd[index]
		if ch == '\r' {
			index += 2 // Move past '\r\n'
			break
		}
		if !(ch >= '0' && ch <= '9') {
			return nil, errors.Err{Msg: fmt.Sprintf("Malformed error from Deserializer for %s\n", rawCmd), File: "resp/deserializer.go", Line: 73}
		}

		numSegments = 10*numSegments + int(ch-48)
		index++
	}

	// Read each segment based on the given number of segments
	for numSegments > 0 {
		length, nextIdx, err := getSegmentLength(rawCmd, index)
		if err != nil {
			return nil, err
		}
		seg, nextIdx, err1 := getSegment(rawCmd, int(nextIdx), int(length))
		if err1 != nil {
			return nil, err1
		}

		index = int(nextIdx)
		segments = append(segments, string(seg))
		numSegments--
	}

	if index != len(rawCmd) {
		return nil, errors.Err{Msg: fmt.Sprintf("Malformed error from Deserializer for %s\n", rawCmd), File: "resp/deserializer.go", Line: 97}
	}
	return segments, nil
}
