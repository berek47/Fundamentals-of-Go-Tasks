package main

import (
	"fmt"
	"library_management/controllers"
	"library_management/services"
	"os"
)

var libraryService *services.LibraryService

func main() {
	// Initialize service
	libraryService = services.NewLibraryService()

	// Main menu loop
	for {
		// Display menu
		fmt.Println("Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Select an option: ")

		var option int
		_, err := fmt.Scanf("%d", &option) 
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			var dummy string
			fmt.Scanln(&dummy)
			continue
		}
		switch option {
		case 1:
			controllers.AddBook(libraryService)
		case 2:
			controllers.RemoveBook(libraryService)
		case 3:
			controllers.BorrowBook(libraryService)
		case 4:
			controllers.ReturnBook(libraryService)
		case 5:
			controllers.ListAvailableBooks(libraryService)
		case 6:
			controllers.ListBorrowedBooks(libraryService)
		case 7:
			fmt.Println("Exiting Library Management System.")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
