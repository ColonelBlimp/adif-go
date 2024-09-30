package adif

import "unicode"

func isNthRuneFromRightEqual(s string, n int, char rune) bool {
	runes := []rune(s) // Convert the string to a slice of runes
	if n > len(runes) || n <= 0 {
		return false // Out of bounds, return false
	}
	return runes[len(runes)-n] == char
}

func isAllDigits(s string) bool {
	if len(s) == 0 {
		return false // Consider empty string as not all digits
	}
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
