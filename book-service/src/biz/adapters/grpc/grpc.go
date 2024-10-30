package grpc

import (
	"context"

	"github.com/huynhminhtruong/go-store-services/book-service/src/biz/application/core/domain"
	"github.com/huynhminhtruong/go-store-services/book-service/src/services/book"
)

// Create import reference type is [go_package_name].[type_struct_name]
func (a Adapter) Create(ctx context.Context, request *book.CreateBookRequest) (*book.CreateBookResponse, error) {
	// business domain supplies models
	newBook := domain.NewBook(request.Title, request.Author, request.PublishYear)
	// gRPC adapter is using APIPort interface to call business logic inside application
	result, err := a.api.InsertBook(newBook)
	if err != nil {
		return nil, err
	}
	return &book.CreateBookResponse{BookId: result.BookID}, nil
}

// GetBook by book_id
func (a Adapter) GetBook(ctx context.Context, request *book.GetBookRequest) (*book.GetBookResponse, error) {
	bookID := request.BookId
	result, err := a.api.GetBook(bookID)
	if err != nil {
		return nil, err
	}
	bookRes := &book.GetBookResponse{
		Title:       result.Title,
		Author:      result.Author,
		PublishYear: result.PublishYear,
	}
	return bookRes, nil
}

// ListBooks TODO get list of books
func (a Adapter) ListBooks(ctx context.Context, request *book.ListBooksRequest) (*book.ListBooksResponse, error) {
	return &book.ListBooksResponse{}, nil
}
