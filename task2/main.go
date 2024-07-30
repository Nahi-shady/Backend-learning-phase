package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	frequency_counter "task2/frequency-counter"
	palindrome_checker "task2/palindrome-checker"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Nahom's/Jhon's String Processor")
	fmt.Println("-------------------------------")

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Check if a string is a palindrome")
		fmt.Println("2. Calculate word frequency")
		fmt.Println("3. Exit")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter a string to check for palindrome: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if palindrome_checker.IsPalindrome(input) {
				fmt.Println("The string is a palindrome!")
			} else {
				fmt.Println("The string is not a palindrome.")
			}

		case "2":
			fmt.Print("Enter a text to analyze word frequency: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)

			frequencies := frequency_counter.WordFrequency(text)
			fmt.Println("Word Frequencies:")
			for word, count := range frequencies {
				fmt.Printf("%s: %d\n", word, count)
			}

		case "3":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
