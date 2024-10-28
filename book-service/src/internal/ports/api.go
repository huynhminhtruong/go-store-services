package ports

import "github.com/huynhminhtruong/go-store-services/book-service/src/src/internal/application/core/domain"

type APIPort interface {
	InsertBook(book domain.Book) (domain.Book, error)
}

type DBPort interface {
	Get(id string) error
	Save(book *domain.Book) error
}