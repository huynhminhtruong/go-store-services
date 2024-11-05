package ports

import "github.com/huynhminhtruong/go-store-services/book-service/src/biz/application/core/domain"

type APIPort interface {
	InsertBook(book domain.Book) (domain.Book, error)
	GetBook(id int64) (domain.Book, error)
	ListBooks() ([]domain.Book, error)
}

type DBPort interface {
	Get(id int64) (domain.Book, error)
	Save(book *domain.Book) *domain.CreateBookResponse
	GetListBooks() ([]domain.Book, error)
}
