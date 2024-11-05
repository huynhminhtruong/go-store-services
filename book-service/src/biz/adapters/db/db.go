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

type CreateBookResponse struct {
	BookID       int64 `json:"book_id"`
	RowsAffected int64 `json:"rows_affected"`
	ErrorMessage error `json:"error_message"`
}

func (a Adapter) Save(book *domain.Book) *CreateBookResponse {
	newBook := domain.NewBook(book.Title, book.Author, book.PublishYear)
	bookResponse := CreateBookResponse{}
	result := a.db.Create(&newBook)

	if result.Error != nil {
		bookResponse.BookID = -1
		bookResponse.RowsAffected = 0
		bookResponse.ErrorMessage = result.Error
	} else {
		bookResponse.BookID = newBook.ID
		bookResponse.RowsAffected = result.RowsAffected
		bookResponse.ErrorMessage = nil
	}

	return &bookResponse
}

func (a Adapter) GetListBooks() ([]domain.Book, error) {
	var books []domain.Book

	if err := a.db.Find(&books).Error; err != nil {
		return []domain.Book{}, err
	}

	return books, nil
}
