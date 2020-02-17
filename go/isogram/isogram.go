package isogram

import "unicode"

func IsIsogram(input string) bool {
	dictLetter := map[rune]bool{}
	for _, letter := range input {
		if !unicode.IsLetter(letter) {
			continue
		}
		lowerLetter := unicode.ToLower(letter)
		if dictLetter[lowerLetter] {
			return false
		}
		dictLetter[lowerLetter] = true
	}
	return true
}
