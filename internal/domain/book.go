package domain

import "errors"

var (
	ErrBookNotFound = errors.New("book not found")
	ErrInvalidInput = errors.New("invalid input")
)

type Book struct {
	ID     string
	Title  string
	Author string
	Year   int
}

func NewBook(id, title, author string, year int) (*Book, error) {
	if title == "" || author == "" {
		return nil, ErrInvalidInput
	}
	if year < 0 || year > 2100 {
		return nil, ErrInvalidInput
	}

	return &Book{
		ID:     id,
		Title:  title,
		Author: author,
		Year:   year,
	}, nil
}