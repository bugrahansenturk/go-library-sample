package services

import (
	"errors"
	"sync"

	domain "library-sample/domains"
)

type BookService struct {
	books   map[int]domain.Book
	idMutex sync.Mutex
	nextID  int
}

func NewBookService() *BookService {
	service := &BookService{
		books:  make(map[int]domain.Book),
		nextID: 1,
	}

	// Add initial books
	service.AddBook(domain.Book{
		Title:  "The Catcher in the Rye",
		Author: "J.D. Salinger",
	})
	service.AddBook(domain.Book{
		Title:  "To Kill a Mockingbird",
		Author: "Harper Lee",
	})
	service.AddBook(domain.Book{
		Title:  "1984",
		Author: "George Orwell",
	})

	return service
}

func (s *BookService) ListBooks() []domain.Book {
	books := make([]domain.Book, 0, len(s.books))
	for _, book := range s.books {
		books = append(books, book)
	}
	return books
}

func (s *BookService) GetBookByID(id int) (domain.Book, error) {
	book, ok := s.books[id]
	if !ok {
		return domain.Book{}, errors.New("book not found")
	}
	return book, nil
}

func (s *BookService) AddBook(book domain.Book) (domain.Book, error) {
	s.idMutex.Lock()
	book.ID = s.nextID
	s.nextID++
	s.idMutex.Unlock()

	s.books[book.ID] = book
	return book, nil
}

func (s *BookService) UpdateBook(updatedBook domain.Book) error {
	_, ok := s.books[updatedBook.ID]
	if !ok {
		return errors.New("book not found")
	}
	s.books[updatedBook.ID] = updatedBook
	return nil
}

func (s *BookService) DeleteBook(id int) error {
	_, ok := s.books[id]
	if !ok {
		return errors.New("book not found")
	}
	delete(s.books, id)
	return nil
}
