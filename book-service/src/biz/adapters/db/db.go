package db

import (
	"fmt"

	"github.com/huynhminhtruong/go-store-services/book-service/src/biz/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model  // Adds entity metadata such as ID to struct
	BookID      int64
	Title       string
	Author      string
	PublishYear string
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceURL string) (*Adapter, error) {
	db, openErr := gorm.Open(postgres.Open(dataSourceURL), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	// create database table
	err := db.AutoMigrate(&Book{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) Get(bookID int64) (domain.Book, error) {
	var book domain.Book

	err := a.db.First(&book, "book_id = ?", bookID).Error
	if err != nil {
		return domain.Book{}, err
	}

	return book, nil
}

func (a Adapter) Save(book *domain.Book) error { return nil }

func (a Adapter) GetListBooks() ([]domain.Book, error) {
	var books []domain.Book

	if err := a.db.Find(&books).Error; err != nil {
		return []domain.Book{}, err
	}

	return books, nil
}
