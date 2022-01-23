package repositories

import (
	"context"
	"net/url"
	"quiz-3/models"
)

type BookRepository interface {
	GetBooks(ctx context.Context, query url.Values) ([]models.Book, error)
	GetBookById(ctx context.Context, id int) (models.Book, error)
	CreateBook(ctx context.Context, book models.Book) (models.Book, error)
	UpdateBook(ctx context.Context, book models.Book) (models.Book, error)
	DeleteBook(ctx context.Context, id int) error
}