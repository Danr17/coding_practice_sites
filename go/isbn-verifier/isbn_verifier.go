package isbn

import (
	"strings"
	"unicode"
)

//IsValidISBN checks if the isbn is valid
func IsValidISBN(isbn string) bool {

	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i, s := range isbn {
		if s == 'X' && i == 9 {
			sum += 10
			continue
		}
		if !(unicode.IsDigit(s)) {
			return false
		}
		sum += int(s-'0') * (10 - i)
	}

	return sum%11 == 0
}
