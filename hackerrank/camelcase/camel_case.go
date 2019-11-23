package camelcase

import (
	"unicode"
)

// CamelCaseWordCounter for given string s the function will return the number
// words taking into account that each uppercase letter will count as a new word
func CamelCaseWordCounter(s string) (count int32) {
	count = 1
	for _, r := range s {
		if unicode.IsUpper(r) {
			count++
		}
	}
	return count
}
