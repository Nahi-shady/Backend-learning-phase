# Library Management System

The **Library Management System** is a simple command-line application written in Go that allows you to manage a library's collection of books and members. The system supports various functionalities, such as adding and removing books, borrowing and returning books, listing available and borrowed books, and managing library members.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Code Structure](#code-structure)

## Features

The Library Management System provides the following features:

1. **Add New Book**: Allows adding a new book to the library collection.
2. **Remove Book**: Allows removing a book from the library by its ID.
3. **Borrow Book**: Allows a registered member to borrow an available book.
4. **Return Book**: Allows a member to return a borrowed book.
5. **List Available Books**: Displays a list of all available books in the library.
6. **List Borrowed Books**: Displays a list of books borrowed by a specific member.
7. **Add New Member**: Adds a new member to the library's member list.

## Requirements

- **Go**: Version 1.16 or later

You can install Go from the official website: [golang.org](https://golang.org/doc/install).

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/library-management-system.git
```

## Usage

After running the application, you will be presented with a menu to interact with the library system:

--------------------------------------
Library Management System

--------------------------------------

1. Add New Book
2. Remove Book
3. Borrow Book
4. Return Book
5. List Available Books
6. List Borrowed Books
7. Add A New Member
8. Exit
Select an option:

## Code structure

The codebase is structured into several packages to separate concerns and organize the functionality:

models: Contains data structures for books and members
services: Implements the library operations.
controllers: Handles user input and orchestrates actions.
main.go: Entry point for the application, with the command-line interface logic.
Here's an overview of the main components:
