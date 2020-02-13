package luhn

import (
	"strings"
	"unicode"
)

func Valid(input string) bool {
	cleanInput := strings.ReplaceAll(input, " ", "")

	if len(cleanInput) < 2 {
		return false
	}

	checkSum := 0
	checkCurrent := false
	for i := len(cleanInput) - 1; i >= 0; i-- {

		current := rune(cleanInput[i])
		if unicode.IsDigit(current) {
			if checkCurrent {
				checkSum += ExtractLuhnValueFromDigit(current)
			} else {
				checkSum += int(current - '0')
			}
		} else {
			return false
		}
		checkCurrent = !checkCurrent
	}

	return checkSum%10 == 0
}

func ExtractLuhnValueFromDigit(digit rune) int {
	n := int(digit - '0')

	if n*2 > 9 {
		n = n*2 - 9
	} else {
		n = n * 2
	}

	return n
}
