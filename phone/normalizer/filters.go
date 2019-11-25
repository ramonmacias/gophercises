package normalizer

import "strings"

// FilterFuncs will return an slice of filters functions
func FilterFuncs() (filters []func(string) string) {
	filters = append(filters, FilterParenthesis())
	filters = append(filters, FilterHyphens())
	return filters
}

// FilterParenthesis will return a function that filters parenthesis
func FilterParenthesis() func(string) string {
	return func(phoneNumber string) string {
		return strings.ReplaceAll(strings.ReplaceAll(phoneNumber, "(", ""), ")", "")
	}
}

// FilterHyphens will return a function that filters hyphens
func FilterHyphens() func(string) string {
	return func(phoneNumber string) string {
		return strings.ReplaceAll(phoneNumber, "-", "")
	}
}
