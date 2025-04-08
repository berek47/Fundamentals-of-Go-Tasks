package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

type LibraryController struct {
	service *services.LibraryService
}

func NewLibraryController() *LibraryController {
	return &LibraryController{
		service: services.NewLibraryService(),
	}
}

func (c *LibraryController) AddBook() {
	var title, author string
	fmt.Print("Enter book title: ")
	fmt.Scan(&title)
	fmt.Print("Enter author name: ")
	fmt.Scan(&author)

	book := models.Book{
		Title:  title,
		Author: author,
		Status: "Available",
	}
	c.service.AddBook(book)
	fmt.Println("Book added successfully!")
}

func (c *LibraryController) RemoveBook() {
	var id int
	fmt.Print("Enter book ID to remove: ")
	fmt.Scan(&id)
	err := c.service.RemoveBook(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book removed successfully!")
	}
}

func (c *LibraryController) BorrowBook() {
	var bookID, memberID int
	fmt.Print("Enter book ID to borrow: ")
	fmt.Scan(&bookID)
	fmt.Print("Enter member ID: ")
	fmt.Scan(&memberID)
	err := c.service.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func (c *LibraryController) ReturnBook() {
	var bookID, memberID int
	fmt.Print("Enter book ID to return: ")
	fmt.Scan(&bookID)
	fmt.Print("Enter member ID: ")
	fmt.Scan(&memberID)
	err := c.service.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func (c *LibraryController) ListAvailableBooks() {
	books := c.service.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) ListBorrowedBooks() {
	var memberID int
	fmt.Print("Enter member ID: ")
	fmt.Scan(&memberID)
	books := c.service.ListBorrowedBooks(memberID)
	if len(books) > 0 {
		fmt.Println("Borrowed Books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	} else {
		fmt.Println("No borrowed books found.")
	}
}
