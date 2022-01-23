package repositories

import (
	"context"
	"database/sql"
	"quiz-3/models"
	"time"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (repository *CategoryRepositoryImpl) GetCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	script := "SELECT * FROM categories"

	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return categories, err
	}

	defer rows.Close()

	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return categories, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (repository *CategoryRepositoryImpl) GetCategoryById(ctx context.Context, id int) (models.Category, error) {
	var category models.Category
	script := "SELECT * FROM categories WHERE id = ?"

	row := repository.DB.QueryRowContext(ctx, script, id)
	err := row.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) CreateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	script := "INSERT INTO categories (name, created_at, updated_at) VALUES (?, ?, ?)"

	now := time.Now().Format("2006-01-02 15:04:05")

	result, err := repository.DB.ExecContext(ctx, script, category.Name, now, now)
	if err != nil {
		return category, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return category, err
	}

	category.ID = int(id)

	return category, nil
}

func (repository *CategoryRepositoryImpl) UpdateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	script := "UPDATE categories SET name = ?, updated_at = ? WHERE id = ?"

	now := time.Now().Format("2006-01-02 15:04:05")

	result, err := repository.DB.ExecContext(ctx, script, category.Name, now, category.ID)
	if err != nil {
		return category, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return category, err
	}

	if rows == 0 {
		return category, sql.ErrNoRows
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) DeleteCategory(ctx context.Context, id int) error {
	script := "DELETE FROM categories WHERE id = ?"

	result, err := repository.DB.ExecContext(ctx, script, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
