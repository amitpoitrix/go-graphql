package repository

import (
	"sync"
	"github.com/amitpoitrix/go-graphql/internal/domain"
)

type BookRepository interface {
	FindAll() ([]*domain.Book, error)
	FindByID(id string) (*domain.Book, error)
	Create(book *domain.Book) (*domain.Book, error)
	Update(book *domain.Book) (*domain.Book, error)
	Delete(id string) error
}

type bookRepository struct {
	mu    sync.RWMutex
	books map[string]*domain.Book
}

func NewBookRepository() BookRepository {
	repo := &bookRepository{
		books: make(map[string]*domain.Book),
	}
	// Seed data
	repo.seedData()
	return repo
}

func (r *bookRepository) seedData() {
	books := []*domain.Book{
		{ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925},
		{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Year: 1960},
		{ID: "3", Title: "1984", Author: "George Orwell", Year: 1949},
	}
	for _, book := range books {
		r.books[book.ID] = book
	}
}

func (r *bookRepository) FindAll() ([]*domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	books := make([]*domain.Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}
	return books, nil
}

func (r *bookRepository) FindByID(id string) (*domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	book, exists := r.books[id]
	if !exists {
		return nil, domain.ErrBookNotFound
	}
	return book, nil
}

func (r *bookRepository) Create(book *domain.Book) (*domain.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.books[book.ID] = book
	return book, nil
}

func (r *bookRepository) Update(book *domain.Book) (*domain.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.books[book.ID]; !exists {
		return nil, domain.ErrBookNotFound
	}
	r.books[book.ID] = book
	return book, nil
}

func (r *bookRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.books[id]; !exists {
		return domain.ErrBookNotFound
	}
	delete(r.books, id)
	return nil
}