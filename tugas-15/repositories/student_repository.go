package repositories

import (
	"context"
	Models "tugas-15/models"
)

type StudentRepository interface {
	GetAll(ctx context.Context) ([]Models.Student, error)
	Get(ctx context.Context, id int32) (Models.Student, error)
	Insert(ctx context.Context, student Models.Student) (Models.Student, error)
	Update(ctx context.Context, student Models.Student, id int32) (Models.Student, error)
	Delete(ctx context.Context, id int32) error
}