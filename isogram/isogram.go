/*
Package isogram contains the solution to the Exercism:Isogram problem.
*/
package isogram

import (
	"regexp"
	"strings"
)

// IsIsogram inspects a string and determines if it is a word or phrase
// without a repeating letter. However, spaces and hyphens are allowed
// to appear multiple times.
func IsIsogram(input string) bool {
	letters := make(map[rune]rune)

	for _, letter := range CleanInput(input) {
		if _, ok := letters[letter]; ok {
			return false
		} else {
			letters[letter] = letter
		}
	}
	return true
}

// CleanInput removes whitespace and hypen characters first, and then switches
// everything to lower case.
func CleanInput(input string) string {
	// Match 1 or more whitespace characters OR 1 or more hyphen characters
	r := regexp.MustCompile(`\s+|-+`)
	return strings.ToLower(r.ReplaceAllString(input, ""))
}
