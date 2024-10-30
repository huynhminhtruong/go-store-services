package api

import (
	"github.com/huynhminhtruong/go-store-services/book-service/src/biz/application/core/domain"
	"github.com/huynhminhtruong/go-store-services/book-service/src/biz/ports"
)

type Application struct {
	db ports.DBPort // database interface
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

// InsertBook method of ports.APIPort interface
func (a Application) InsertBook(book domain.Book) (domain.Book, error) {
	if err := a.db.Save(&book); err != nil {
		return domain.Book{}, err
	}

	return book, nil
}

func (a Application) GetBook(id int64) (domain.Book, error) {
	book, err := a.db.Get(id)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (a Application) ListBooks() ([]domain.Book, error) {
	return []domain.Book{}, nil
}
