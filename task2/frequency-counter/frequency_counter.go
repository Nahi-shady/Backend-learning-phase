// frequency_counter.go
package frequency_counter

import (
	"strings"
	"unicode"
)

func WordFrequency(input string) map[string]int {
	lowercaseInput := strings.ToLower(input)

	var builder strings.Builder
	for _, char := range lowercaseInput {
		if unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsSpace(char) {
			builder.WriteRune(char)
		}
	}

	words := strings.Fields(builder.String())

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}
