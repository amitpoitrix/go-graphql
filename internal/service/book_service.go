package service

import (
	"fmt"
	"github.com/amitpoitrix/go-graphql/internal/domain"
	"github.com/amitpoitrix/go-graphql/internal/repository"
)

type BookService interface {
	GetAllBooks() ([]*domain.Book, error)
	GetBookByID(id string) (*domain.Book, error)
	CreateBook(title, author string, year int) (*domain.Book, error)
	UpdateBook(id, title, author string, year int) (*domain.Book, error)
	DeleteBook(id string) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() ([]*domain.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetBookByID(id string) (*domain.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) CreateBook(title, author string, year int) (*domain.Book, error) {
	books, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	id := fmt.Sprintf("%d", len(books)+1)
	book, err := domain.NewBook(id, title, author, year)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(book)
}

func (s *bookService) UpdateBook(id, title, author string, year int) (*domain.Book, error) {
	book, err := domain.NewBook(id, title, author, year)
	if err != nil {
		return nil, err
	}
	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id string) error {
	return s.repo.Delete(id)
}