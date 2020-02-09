// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb makes a proverb from a list of strings
func Proverb(rhyme []string) []string {

	proverb := []string{}

	if len(rhyme) == 1 {
		proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}

	if len(rhyme) > 1 {

		for i := 0; i < len(rhyme)-1; i++ {
			proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
		proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}

	return proverb
}
