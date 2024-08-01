package services

import (
	"fmt"
	"library_management/models"
	"time"
)

type LibraryManager interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks()
	ListBorrowedBooks(memberID int)
	AddMember(member models.Member) error
}

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func display_books(books map[int]models.Book) {
	fmt.Println("ID ", "|Title ", "|Author ", "|Status")
	fmt.Println("----------------------------------------------------------------------------")
	for _, book := range books {
		fmt.Printf("%d: %v |%v |%v\n", book.ID, book.Title, book.Author, book.Status)
		fmt.Println("-------------------------------------------------------------------")
	}
}

func display_members(members map[int]models.Member) {
	fmt.Println("ID  ", "Name  ")
	for _, member := range members {
		fmt.Printf("%d: %s\n", member.ID, member.Name)
	}
}

func (l *Library) AddBook(book models.Book) error {
	if _, exists := l.books[book.ID]; exists {
		return fmt.Errorf("book with ID %d already exists", book.ID)
	}
	l.books[book.ID] = book
	return nil
}

func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.books[bookID]; !exists {
		return fmt.Errorf("book with ID %d does not exist", bookID)
	}
	delete(l.books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return fmt.Errorf("book with ID %d not available", bookID)
	}

	if book.Status != "Available" {
		return fmt.Errorf("book with ID %d is not available for borrowing", bookID)
	}

	member, memberExists := l.members[memberID]
	if !memberExists {
		return fmt.Errorf("member with ID %d not found", memberID)
	}

	if member.Borrowed == nil {
		member.Borrowed = make(map[int]models.Book)
	}

	book.Status = "Borrowed"
	l.books[bookID] = book
	member.Borrowed[bookID] = book
	l.members[memberID] = member

	fmt.Printf("Member %s successfully borrowed book %s", member.Name, book.Title)
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return fmt.Errorf("book with ID %d not found in the library", bookID)
	}

	member, memberExists := l.members[memberID]
	if !memberExists {
		return fmt.Errorf("member with ID %d not found", memberID)
	}

	if _, borrowed := member.Borrowed[bookID]; !borrowed {
		return fmt.Errorf("book with ID %d was not borrowed by member with ID %d", bookID, memberID)
	}

	book.Status = "Available"
	l.books[bookID] = book
	delete(member.Borrowed, bookID)
	l.members[memberID] = member
	return nil
}

func (l *Library) ListAvailableBooks() {
	display_books(l.books)
}

func (l *Library) ListBorrowedBooks(memberID int) {
	member, exists := l.members[memberID]
	if !exists {
		fmt.Println("Member with that ID doesn't exist!")
	}

	display_books(member.Borrowed)
}

func (l *Library) AddMember(member models.Member) error {
	if _, exists := l.members[member.ID]; exists {
		return fmt.Errorf("member with ID %d already exists", member.ID)
	}
	l.members[member.ID] = member
	display_members(l.members)
	time.Sleep(time.Duration(2) * time.Second)

	return nil
}
