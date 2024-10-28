package domain

/*
	Business logic here
*/

type Book struct {
	ID          int64  `json:"id"`
	BookID      int64  `json:"book_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishYear string `json:"publish_year"`
}

func NewBook(title, author, publishYear string) Book {
	return Book{
		Title:       title,
		Author:      author,
		PublishYear: publishYear,
	}
}
