package repositories

import (
	"context"
	"quiz-3/models"
)

type CategoryRepository interface {
	GetCategories(ctx context.Context) ([]models.Category, error)
	GetCategoryById(ctx context.Context, id int) (models.Category, error)
	CreateCategory(ctx context.Context, category models.Category) (models.Category, error)
	UpdateCategory(ctx context.Context, category models.Category) (models.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}
