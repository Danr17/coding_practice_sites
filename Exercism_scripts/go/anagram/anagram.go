package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	orderedSubject := SplitandSort(subject)
	allPositive := []string{}

Outer:
	for _, candidate := range candidates {
		orderedCandidate := SplitandSort(candidate)
		if len(orderedCandidate) != len(orderedSubject) || strings.ToLower(subject) == strings.ToLower(candidate) {
			continue
		}
		for i, v := range orderedCandidate {
			if v != orderedSubject[i] {
				continue Outer
			}
		}
		allPositive = append(allPositive, candidate)
	}
	return allPositive
}

func SplitandSort(item string) []string {
	splitted := make([]string, len(item)+1)
	for i, s := range item {
		splitted[i] = strings.ToLower(string(s))

	}
	sort.Strings(splitted)

	return splitted
}
