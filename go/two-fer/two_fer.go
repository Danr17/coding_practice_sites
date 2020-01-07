// Package twofer implements the two for one
package twofer

import "fmt"

// ShareWith takes a string and return a concatenated text
func ShareWith(name string) string {

	var text string

	if name == "" {
		text = fmt.Sprintln("One for you, one for me.")
		return text
	}

	text = fmt.Sprintf("One for %s, one for me.", name)

	return text
}
