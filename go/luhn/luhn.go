package luhn

import (
	"unicode"
)

func Valid(input string) bool {

	sumDouble1 := 0
	sumSimple1 := 0
	sumDouble2 := 0
	sumSimple2 := 0
	counter := 0
	for _, r := range input {

		if unicode.IsDigit(r) {
			val := int(r - '0')
			doubleval := (val * 2)
			if doubleval > 9 {
				doubleval = doubleval - 9

			}
			if counter%2 == 0 {
				sumDouble1 += doubleval
				sumSimple1 += val
			} else {
				sumDouble2 += doubleval
				sumSimple2 += val
			}
			counter++
			continue
		}

		if unicode.IsSpace(r) {
			continue
		}
		return false
	}

	if counter < 2 {
		return false
	}

	if counter%2 == 0 {
		return ((sumDouble1 + sumSimple2) % 10) == 0
	}
	return ((sumDouble2 + sumSimple1) % 10) == 0
}
