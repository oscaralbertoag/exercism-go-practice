package bob

import (
	"regexp"
	"strings"
	"unicode"
)

func Hey(remark string) string {
	result := "Whatever."
	if IsAllWhitespace(remark) {
		result = "Fine. Be that way!"
	} else if IsQuestion(remark) {
		if ContainsAnyLetter(remark) && AllLettersAreUpperCase(remark) {
			result = "Calm down, I know what I'm doing!"
		} else {
			result = "Sure."
		}
	} else if ContainsAnyLetter(remark) && AllLettersAreUpperCase(remark) {
		result = "Whoa, chill out!"
	}
	return result
}

func IsQuestion(s string) bool {
	trimmed := strings.TrimRight(s, " ")
	if match, _ := regexp.MatchString(`\?$`, trimmed); !match {
		return false
	}
	return true
}

func AllLettersAreUpperCase(s string) bool {
	for _, r := range s {
      if !unicode.IsUpper(r) && unicode.IsLetter(r) {
          return false
      }
  }
	return true
}

func ContainsAnyLetter(s string) bool {
	if match, _ := regexp.MatchString(`[a-zA-Z]+`, s); match {
		return true
	}
	return false
}

func IsAllWhitespace(s string) bool {
	for _, r := range s {
		if match, _ := regexp.MatchString(`\s+`, string(r)); !match {
			return false
		}
	}
	return true
}
