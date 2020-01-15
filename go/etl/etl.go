package etl

import "strings"

//Transform is an ETL for Scrabble scores
func Transform(input map[int][]string) map[string]int {
	value := make(map[string]int)

	for i, x := range input {
		for _, z := range x {
			value[strings.ToLower(z)] = i
		}
	}
	return value
}
