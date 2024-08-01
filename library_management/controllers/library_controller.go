package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
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

type LibraryController struct {
	service services.LibraryManager
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{service: service}
}

func (lc *LibraryController) AddNewBook() {
	title := readStringInput("Please enter the title of the book: ")
	author := readStringInput("Name of the author: ")
	book := models.Book{ID: BookIDGenerator(), Title: title, Author: author, Status: "Available"}

	err := lc.service.AddBook(book)
	if err != nil {
		fmt.Println("Error adding book:", err)
		return
	}

	fmt.Println("Book added successfully.")
}
func (lc *LibraryController) RemoveBook() {
	id := readIntInput("Please enter the id of the book: ")
	err := lc.service.RemoveBook(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Book is deleted successfully!")
}

func (lc *LibraryController) BorrowBook() {
	bookID := readIntInput("Enter the Book ID to borrow: ")
	memberID := readIntInput("Enter the Member ID: ")

	err := lc.service.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
		return
	}

	fmt.Println("Book borrowed successfully.")
}

func (lc *LibraryController) ReturnBook() {
	bookID := readIntInput("Enter the Book ID to return: ")
	memberID := readIntInput("Enter the Member ID: ")

	err := lc.service.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error returning book:", err)
		return
	}

	fmt.Println("Book returned successfully.")
}

func (lc *LibraryController) ListAvailableBooks() {
	lc.service.ListAvailableBooks()
}

func (lc *LibraryController) ListBorrowedBooks() {
	memberID := readIntInput("Enter the Member ID to view borrowed books: ")

	lc.service.ListBorrowedBooks(memberID)
}

func (lc *LibraryController) AddMember() {
	memberID := MemberIDGenerator()
	memberName := readStringInput("Enter the Member Name: ")

	member := models.Member{ID: memberID, Name: memberName}

	err := lc.service.AddMember(member)
	if err != nil {
		fmt.Println("Error adding book:", err)
		return
	}

}

var bookIDCounter = 1
var memberIDCounter = 1

func BookIDGenerator() int {
	id := bookIDCounter
	bookIDCounter++
	return id
}
func MemberIDGenerator() int {
	id := memberIDCounter
	memberIDCounter++
	return id
}
