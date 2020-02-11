package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	result := ""
	words := ExtractWords(s)
	for _, word := range words {
		result += ExtractFirstLetter(word)
	}
	return result
}

func ExtractFirstLetter(s string) string {
	first := ""
	for _, r := range s {
		if unicode.IsLetter(r) {
			first = strings.ToUpper(string(r))
			break
		}
	}
	return first
}

func ExtractWords(s string) []string {
	var words []string
	for _, chunk := range strings.Split(s, " ") {
		for _, token := range strings.Split(chunk, "-") {
			words = append(words, token)
		}
	}
	return words
}
