package secret

import (
	"strconv"
)

var actions map[int]string = map[int]string{
	0: "wink",
	1: "double blink",
	2: "close your eyes",
	3: "jump",
}

func Handshake(code uint) []string {

	var binStr string = strconv.FormatInt(int64(code), 2)
	result := []string{}
	idx := 0

	for i := len(binStr); i > 0; i-- {
		if i == 1 && len(binStr) == 5 {
			if string(binStr[i-1]) == "1" {
				result = reverse(result)
				continue
			}
		}
		if string(binStr[i-1]) == "1" {
			result = append(result, actions[idx])
		}
		idx++
	}
	return result
}

func reverse(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
