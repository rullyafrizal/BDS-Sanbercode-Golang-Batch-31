package repositories

import (
	"context"
	"tugas-15/models"
)

type GradeRepository interface {
	GetAll(ctx context.Context) ([]models.Grade, error)
	Get(ctx context.Context, id int32) (models.Grade, error)
	Insert(ctx context.Context, grade models.Grade) (models.Grade, error)
	Update(ctx context.Context, grade models.Grade, id int32) (models.Grade, error)
	Delete(ctx context.Context, id int32) error
}