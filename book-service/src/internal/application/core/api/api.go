package api

import (
	"github.com/huynhminhtruong/go-store-services/book-service/src/internal/application/core/domain"
	"github.com/huynhminhtruong/go-store-services/book-service/src/internal/ports"
)

type Application struct {
	db ports.DBPort // database interface
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

// method of ports.APIPort interface
func (a Application) InsertBook(book domain.Book) (domain.Book, error) {
	if err := a.db.Save(&book); err != nil {
		return domain.Book{}, err
	}

	return book, nil
}
