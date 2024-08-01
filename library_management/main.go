package main

import (
	"bufio"
	"fmt"
	"library_management/controllers"
	"library_management/services"
	"os"
	"strconv"
	"strings"
	"time"
)

func readIntInput(prompt string) int {
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
		return number
	}
}

func main() {
	library := services.NewLibrary()
	controller := controllers.NewLibraryController(library)

	for {
		fmt.Println("--------------------------------------")
		fmt.Println("\nLibrary Management System")
		fmt.Println("--------------------------------------")
		fmt.Println("1. Add New Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Add A New Member")
		fmt.Println("8. Exit")
		time.Sleep(time.Duration(2) * time.Second)

		choice := readIntInput("Select an option: ")

		switch choice {
		case 1:
			controller.AddNewBook()
		case 2:
			controller.RemoveBook()
		case 3:
			controller.BorrowBook()
		case 4:
			controller.ReturnBook()
		case 5:
			controller.ListAvailableBooks()
		case 6:
			controller.ListBorrowedBooks()
		case 7:
			controller.AddMember()
		case 8:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
		time.Sleep(time.Duration(3) * time.Second)
	}
}
