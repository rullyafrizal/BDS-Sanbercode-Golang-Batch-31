package repositories

import (
	"context"
	"database/sql"
	"net/url"
	"quiz-3/models"
	"strconv"
	"time"
)

type BookRepositoryImpl struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &BookRepositoryImpl{DB: db}
}

func (repository *BookRepositoryImpl) GetBooks(ctx context.Context, query url.Values) ([]models.Book, error) {
	var books []models.Book
	script := "SELECT * FROM books "

	var q []interface{}

	if query.Get("title") != "" {
		script += "WHERE title LIKE ? "
		q = append(q, "%"+query.Get("title")+"%")
	}

	if query.Get("minYear") != "" {
		if len(q) > 0 {
			script += "AND "
		} else {
			script += "WHERE "
		}

		script += "release_year >= ? "
		val, _ := strconv.Atoi(query.Get("minYear"))

		q = append(q, val)
	}

	if query.Get("maxYear") != "" {
		if len(q) > 0 {
			script += "AND "
		} else {
			script += "WHERE "
		}

		script += "release_year <= ? "
		val, _ := strconv.Atoi(query.Get("maxYear"))

		q = append(q, val)
	}

	if query.Get("minPage") != "" {
		if len(q) > 0 {
			script += "AND "
		} else {
			script += "WHERE "
		}

		script += "total_page >= ? "
		val, _ := strconv.Atoi(query.Get("minPage"))

		q = append(q, val)
	}

	if query.Get("maxPage") != "" {
		if len(q) > 0 {
			script += "AND "
		} else {
			script += "WHERE "
		}

		script += "total_page <= ? "
		val, _ := strconv.Atoi(query.Get("maxPage"))

		q = append(q, val)
	}

	if query.Get("sortByTitle") != "" {
		script += "ORDER BY title "
		if query.Get("sortByTitle") == "desc" {
			script += "DESC "
		} else {
			script += "ASC "
		}
	}

	rows, err := repository.DB.QueryContext(ctx, script, q...)
	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID, &book.Title,
			&book.Description, &book.ImageUrl,
			&book.ReleaseYear, &book.Price,
			&book.TotalPage, &book.Thickness,
			&book.CategoryId, &book.CreatedAt, &book.UpdatedAt,
		)

		if err != nil {
			return books, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (repository *BookRepositoryImpl) GetBookById(ctx context.Context, id int) (models.Book, error) {
	var book models.Book
	script := "SELECT * FROM books WHERE id = ?"

	err := repository.DB.QueryRowContext(ctx, script, id).Scan(
		&book.ID, &book.Title,
		&book.Description, &book.ImageUrl,
		&book.ReleaseYear, &book.Price,
		&book.TotalPage, &book.Thickness,
		&book.CategoryId, &book.CreatedAt, &book.UpdatedAt,
	)

	if err != nil {
		return book, err
	}

	return book, nil
}

func (repository *BookRepositoryImpl) CreateBook(ctx context.Context, book models.Book) (models.Book, error) {
	script := "INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	now := time.Now().Format("2006-01-02 15:04:05")

	result, err := repository.DB.ExecContext(ctx, script, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryId, now, now)
	if err != nil {
		return book, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return book, err
	}

	book.ID = int(id)

	return book, nil
}

func (repository *BookRepositoryImpl) UpdateBook(ctx context.Context, book models.Book) (models.Book, error) {
	script := "UPDATE books SET title = ?, description = ?, image_url = ?, release_year = ?, price = ?, total_page = ?, thickness = ?, category_id = ?, updated_at = ? WHERE id = ?"

	now := time.Now().Format("2006-01-02 15:04:05")

	result, err := repository.DB.ExecContext(ctx, script, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryId, now, book.ID)
	if err != nil {
		return book, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return book, err
	}

	if count == 0 {
		return book, sql.ErrNoRows
	}

	return book, nil
}

func (repository *BookRepositoryImpl) DeleteBook(ctx context.Context, id int) error {
	script := "DELETE FROM books WHERE id = ?"

	_, err := repository.DB.ExecContext(ctx, script, id)
	if err != nil {
		return err
	}

	return nil
}
