package grpc

import (
	"context"

	"github.com/huynhminhtruong/go-store-services/book-service/src/biz/application/core/domain"
	"github.com/huynhminhtruong/go-store-services/book-service/src/services/book"
)

// import reference type is [go_package_name].[type_struct_name]
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
