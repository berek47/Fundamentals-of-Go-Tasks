package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

func AddBook(libraryService *services.LibraryService) {
	var title, author string
	fmt.Print("Enter book title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter author name: ")
	fmt.Scanln(&author)

	book := models.Book{Title: title, Author: author}
	libraryService.AddBook(book)
}

func RemoveBook(libraryService *services.LibraryService) {
	var id int
	fmt.Print("Enter book ID to remove: ")
	fmt.Scanln(&id)

	err := libraryService.RemoveBook(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book removed successfully!")
	}
}

func BorrowBook(libraryService *services.LibraryService) {
	var bookID, memberID int
	fmt.Print("Enter book ID to borrow: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)

	err := libraryService.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func ReturnBook(libraryService *services.LibraryService) {
	var bookID, memberID int
	fmt.Print("Enter book ID to return: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)

	err := libraryService.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func ListAvailableBooks(libraryService *services.LibraryService) {
	books := libraryService.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func ListBorrowedBooks(libraryService *services.LibraryService) {
	var memberID int
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)

	books := libraryService.ListBorrowedBooks(memberID)
	if len(books) > 0 {
		fmt.Println("Borrowed Books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	} else {
		fmt.Println("No borrowed books found.")
	}
}
