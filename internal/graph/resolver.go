package graph

import "github.com/amitpoitrix/go-graphql/internal/service"

type Resolver struct {
	bookService service.BookService
}

func NewResolver(bookService service.BookService) *Resolver {
	return &Resolver{
		bookService: bookService,
	}
}