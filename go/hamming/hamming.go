package hamming

import (
	"errors"
	"unicode/utf8"
)

//Distance caculate the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {

	if len([]rune(a)) != len([]rune(b)) {
		return 0, errors.New("strings are not equal")
	}

	var distance int

	for len(a) > 0 {
		rA, sizeA := utf8.DecodeRune([]byte(a))
		rB, sizeB := utf8.DecodeRune([]byte(b))

		if rA != rB {
			distance++
		}
		a = a[sizeA:]
		b = b[sizeB:]
	}

	return distance, nil
}
