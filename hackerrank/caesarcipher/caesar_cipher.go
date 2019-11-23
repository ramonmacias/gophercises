package caesarcipher

import (
	"log"
	"regexp"
	"unicode"
)

// Encrypt this encrypt method don't use anything related with ASCII code, and
// for sure the performance is lower
func Encrypt(s string, k int32) (ret string) {
	var validAlphabet []string
	validAlphabetLower := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	validAlphabetUpper := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	var validID = regexp.MustCompile(`[a-z]|[A-Z]`)
	for _, r := range s {
		if validID.MatchString(string(r)) {
			if unicode.IsUpper(r) {
				validAlphabet = validAlphabetUpper
			} else {
				validAlphabet = validAlphabetLower
			}
			for i, c := range validAlphabet {
				if c == string(r) {
					mov := i + int(k)
					ret += validAlphabet[(mov % len(validAlphabet))]
				}
			}
		} else {
			ret += string(r)
		}
	}
	return ret
}

func EncryptASCII(s string, k int32) (ret string) {
	minLowerRune, maxLowerRune, minUpperRune, maxUpperRune := 'a', 'z', 'A', 'Z'
	log.Println(maxLowerRune, maxUpperRune)

	var validID = regexp.MustCompile(`[a-z]|[A-Z]`)
	for _, r := range s {
		if validID.MatchString(string(r)) {
			mov := r + k
			if unicode.IsUpper(r) {
				ret += string(minUpperRune + (mov % 26))
			} else {
				ret += string(minLowerRune + (mov % 26))
			}
		} else {
			ret += string(r)
		}
	}
	return ret
}
