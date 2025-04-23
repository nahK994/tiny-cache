package cache

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/nahK994/TinyCache/pkg/utils"
)

func createStringItem(value string, expiryTime *time.Time, frequency int) *DataItem {
	bytes := []byte(value)
	return &DataItem{
		DataType:   utils.String,
		Value:      bytes,
		ExpiryTime: expiryTime,
		Frequency:  frequency,
	}
}

func createIntItem(value int, expiryTime *time.Time, frequency int) *DataItem {
	bytes := []byte(strconv.Itoa(value))
	return &DataItem{
		DataType:   utils.Int,
		Value:      bytes,
		ExpiryTime: expiryTime,
		Frequency:  frequency,
	}
}

func createListItem(values []string, expiryTime *time.Time, frequency int) *DataItem {
	bytes, _ := json.Marshal(values)
	return &DataItem{
		DataType:   utils.Array,
		Value:      bytes,
		ExpiryTime: expiryTime,
		Frequency:  frequency,
	}
}

func getList(data []byte) []string {
	if data == nil {
		return []string{}
	}

	var vals []string
	json.Unmarshal(data, &vals)
	return vals
}

// reverseSlice reverses a slice of strings
func reverseSlice(s []string) []string {
	n := len(s)
	result := make([]string, n)
	for i := range s {
		result[n-i-1] = s[i]
	}
	return result
}

// processIdx normalizes index values for LRANGE
func processIdx(vals []string, idx int) int {
	if idx < 0 {
		idx = len(vals) + idx
	}
	if idx < 0 {
		return 0
	}
	if idx >= len(vals) {
		return len(vals) - 1
	}
	return idx
}
