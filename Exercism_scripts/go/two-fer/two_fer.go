// Package twofer implements the two for one
package twofer

import "fmt"

// ShareWith takes a string and return a concatenated text
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}
	text := fmt.Sprintf("One for %s, one for me.", name)

	return text
}
