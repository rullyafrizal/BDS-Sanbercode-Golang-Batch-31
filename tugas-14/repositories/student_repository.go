package repositories

import (
	"context"
	Models "tugas-14/models"
)

type StudentRepository interface {
	GetAll(ctx context.Context) ([]Models.Student, error)
	Insert(ctx context.Context, student Models.Student) (Models.Student, error)
}