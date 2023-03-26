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

func (bookService *BookService) ListBooks() []domain.Book {
	books := make([]domain.Book, 0, len(bookService.books))
	for _, book := range bookService.books {
		books = append(books, book)
	}
	return books
}

func (bookService *BookService) GetBookByID(id int) (domain.Book, error) {
	book, ok := bookService.books[id]
	if !ok {
		return domain.Book{}, errors.New("book not found")
	}
	return book, nil
}

func (bookService *BookService) AddBook(book domain.Book) (domain.Book, error) {
	bookService.idMutex.Lock()
	book.ID = bookService.nextID
	bookService.nextID++
	bookService.idMutex.Unlock()

	bookService.books[book.ID] = book
	return book, nil
}

func (bookService *BookService) UpdateBook(updatedBook domain.Book) error {
	_, ok := bookService.books[updatedBook.ID]
	if !ok {
		return errors.New("book not found")
	}
	bookService.books[updatedBook.ID] = updatedBook
	return nil
}

func (bookService *BookService) DeleteBook(id int) error {
	_, ok := bookService.books[id]
	if !ok {
		return errors.New("book not found")
	}
	delete(bookService.books, id)
	return nil
}
