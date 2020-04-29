package wordcount

import (
	"strings"
	"unicode"
)

// Frequency storage type for phrase words count
type Frequency map[string]int

// WordCount calculates words count in given phrase and returns result prepared in form of map
func WordCount(phrase string) Frequency {
	result := Frequency{}
	wordsarr := strings.FieldsFunc(strings.ToLower(phrase), func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})
	for _, w := range wordsarr {
		if trimw := strings.Trim(w, " '"); trimw != "" {
			result[trimw]++
		}
	}
	return result
}
