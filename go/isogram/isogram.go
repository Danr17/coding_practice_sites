package isogram

import "unicode"

func IsIsogram(input string) bool {
	dictLetter := map[rune]bool{}
	for _, letter := range input {
		if !unicode.IsLetter(letter) {
			continue
		}
		if _, ok := dictLetter[unicode.ToLower(letter)]; ok {
			return false
		}
		dictLetter[unicode.ToLower(letter)] = true
	}
	return true
}
