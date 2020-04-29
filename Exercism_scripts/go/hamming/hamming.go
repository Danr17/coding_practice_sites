package hamming

import (
	"errors"
	"unicode/utf8"
)

//Distance caculate the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {

	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("strings are not equal")
	}

	var distance int

	for len(a) > 0 {
		rA, sizeA := utf8.DecodeRuneInString(a)
		rB, sizeB := utf8.DecodeRuneInString(b)

		if rA != rB {
			distance++
		}
		a = a[sizeA:]
		b = b[sizeB:]
	}

	return distance, nil
}
