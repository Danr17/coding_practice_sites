package acronym

import (
	"strings"
)

// Abbreviate creates acronyms
func Abbreviate(s string) string {
	r := strings.NewReplacer("-", " ", "_", " ")
	cleanS := r.Replace(s)
	split := strings.Split(cleanS, " ")
	output := ""
	for _, word := range split {

		if word == "" {
			continue
		}
		output = output + strings.ToUpper(string(word[0]))
	}

	return output
}
