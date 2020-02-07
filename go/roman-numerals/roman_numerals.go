package romannumerals

import "errors"

//ToRomanNumeral convert digits to Roman numbers
func ToRomanNumeral(input int) (string, error) {
	ArabictoRoman := []struct {
		arabic int
		roman  string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	if input <= 0 || input > 3000 {
		return "", errors.New("the input number it is not in range")
	}

	roman := ""
	for _, number := range ArabictoRoman {
		for input >= number.arabic {
			roman += number.roman
			input -= number.arabic
		}
	}
	return roman, nil

}
