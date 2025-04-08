# Library Management System Documentation

## Features
- **Add Book**: Allows the admin to add a new book to the library.
- **Remove Book**: Allows the admin to remove a book from the library.
- **Borrow Book**: Allows members to borrow available books.
- **Return Book**: Allows members to return borrowed books.
- **List Available Books**: Lists all the books available in the library.
- **List Borrowed Books**: Lists all books borrowed by a member.

## Code Structure
- **main.go**: Entry point of the application that initiates the user interface.
- **controllers/library_controller.go**: Handles user inputs from the console and communicates with the services layer.
- **models/book.go**: Contains the structure and properties of a Book.
- **models/member.go**: Contains the structure and properties of a Member.
- **services/library_service.go**: Contains the business logic for managing books and members.

## Error Handling
- If a book is not found when removing, borrowing, or returning it, an error will be displayed.
- If a member is not found when borrowing or returning a book, an error will be displayed.
