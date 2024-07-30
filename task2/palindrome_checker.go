package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Palindrome Checker")
	fmt.Println("------------------")

	for {
		fmt.Print("Enter a string (or type 'exit' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if isPalindrome(input) {
			fmt.Println("The string is a palindrome!")
		} else {
			fmt.Println("The string is not a palindrome.")
		}
	}
}
