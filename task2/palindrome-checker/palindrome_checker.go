// palindrome_checker.go
package palindrome_checker

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	var normalized []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			normalized = append(normalized, unicode.ToLower(r))
		}
	}

	length := len(normalized)
	for i := 0; i < length/2; i++ {
		if normalized[i] != normalized[length-i-1] {
			return false
		}
	}
	return true
}
