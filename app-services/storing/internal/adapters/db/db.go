package db

import (
	"fmt"

	"github.com/huynhminhtruong/go-store-services/app-services/storing/internal/application/core/domain"
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

	err := db.AutoMigrate(&Book{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) Get(id string) error { return nil }

func (a Adapter) Save(book *domain.Book) error { return nil }
