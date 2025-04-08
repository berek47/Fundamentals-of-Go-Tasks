package services

import (
	"errors"
	"library_management/models"
)

type LibraryService struct {
	books   map[int]models.Book
	members map[int]models.Member
	nextID  int
}

func NewLibraryService() *LibraryService {
	return &LibraryService{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
		nextID:  1,
	}
}

func (s *LibraryService) AddBook(book models.Book) {
	book.ID = s.nextID
	s.books[book.ID] = book
	s.nextID++
}

func (s *LibraryService) RemoveBook(bookID int) error {
	_, exists := s.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	delete(s.books, bookID)
	return nil
}

func (s *LibraryService) BorrowBook(bookID, memberID int) error {
	book, exists := s.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, exists := s.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	s.books[bookID] = book
	s.members[memberID] = member
	return nil
}

func (s *LibraryService) ReturnBook(bookID, memberID int) error {
	book, exists := s.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Available" {
		return errors.New("book was not borrowed")
	}

	member, exists := s.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Available"
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	s.books[bookID] = book
	s.members[memberID] = member
	return nil
}

func (s *LibraryService) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range s.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (s *LibraryService) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := s.members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}
