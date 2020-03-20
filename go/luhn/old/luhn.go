package luhn

import (
	"log"
	"strconv"
	"unicode"
)

//Valid validates the string based on Luhn algorithm
func Valid(input string) bool {
	cleanInput := []int{}
	for _, s := range input {
		if unicode.IsDigit(s) {
			d, err := strconv.Atoi(string(s))
			if err != nil {
				log.Fatal(err)
			}
			cleanInput = append(cleanInput, d)
			continue
		}
		if s != ' ' {
			return false
		}
	}

	if len(cleanInput) < 2 {
		return false
	}

	index := 0
	if len(cleanInput)%2 == 0 {
		index = 1
	}
	sum := 0
	for _, n := range cleanInput {
		index++
		if index%2 == 0 {
			if n*2 > 9 {
				sum = sum + (n*2 - 9)
				continue
			}
			sum = sum + n*2
			continue
		}
		sum = sum + n
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
