package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//RunLengthEncode encode teh string
func RunLengthEncode(input string) string {
	var prevS rune
	count := 1
	output := ""

	if input == "" {
		return ""
	}

	for i, s := range input {
		if i == 0 {
			prevS = s
			continue
		}
		if s == prevS {
			count++
			continue
		}
		if count == 1 {
			output = output + string(prevS)
		} else {
			output = output + strconv.Itoa(count) + string(prevS)
		}
		prevS = s
		count = 1
	}
	if count == 1 {
		output = output + string(prevS)
	} else {
		output = output + strconv.Itoa(count) + string(prevS)
	}

	return output
}

//RunLengthDecode decode the string
func RunLengthDecode(input string) string {
	r := ""
	output := ""

	for _, s := range input {
		if unicode.IsDigit(s) {
			r = r + string(s)
			continue
		}
		if r == "" {
			output = output + string(s)
		} else {
			rep, err := strconv.Atoi(string(r))
			if err != nil {
				fmt.Printf("can't convert rune %v to int\n", r)
			}
			output = output + strings.Repeat(string(s), rep)
			r = ""
		}
	}
	return output
}
