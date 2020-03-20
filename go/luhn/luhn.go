package luhn

import "unicode"

func Valid(input string) bool {
	sum := 0
	counter := 0
	for _, r := range reverse(input) {

		if unicode.IsDigit(r) {
			val := int(r - '0')
			if counter%2 == 1 {
				val = (val * 2)
				if val > 9 {
					val = val - 9
				}
			}
			sum += val
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

	return (sum % 10) == 0
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
