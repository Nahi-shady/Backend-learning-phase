package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readStringInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Input cannot be empty. Please enter a valid string.")
			continue
		}
		return input
	}
}

func readIntInput(prompt string, min, max int) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please enter a valid integer.")
			continue
		}
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}
		if number < min || number > max {
			fmt.Printf("Input out of range. Please enter an integer between %d and %d.\n", min, max)
			continue
		}
		return number
	}
}

func calcAverage(scores map[string]int) float64 {
	var sum int
	for _, score := range scores {
		sum += score
	}
	if len(scores) == 0 {
		return 0
	}
	return float64(sum) / float64(len(scores))
}

func main() {
	fmt.Println("--------------------------")
	fmt.Println("Calculate Your Average Score")
	fmt.Println("--------------------------")

	name := readStringInput("Enter Your Name: ")

	fmt.Printf("Hello, %s!\n", name)

	subjectNumber := readIntInput("How many subjects have you taken: ", 1, 100)

	subjects := make(map[string]int, subjectNumber)

	for i := 0; i < subjectNumber; i++ {
		subject := readStringInput(fmt.Sprintf("Enter Subject %d name: ", i+1))
		score := readIntInput(fmt.Sprintf("Enter %s's score (0-100): ", subject), 0, 100)
		subjects[subject] = score
	}

	average := calcAverage(subjects)

	fmt.Println("--------------------------")
	fmt.Printf("Name: %s\n", name)
	fmt.Println("Subjects and Scores:")
	for subject, score := range subjects {
		fmt.Printf("%s: %d\n", subject, score)
	}
	fmt.Printf("Average Score: %.2f\n", average)
	fmt.Println("--------------------------")
}
