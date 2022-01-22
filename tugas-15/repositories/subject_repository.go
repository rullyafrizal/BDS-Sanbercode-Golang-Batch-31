package repositories

import (
	"context"
	Models "tugas-15/models"
)

type SubjectRepository interface {
	GetAll(ctx context.Context) ([]Models.Subject, error)
	Get(ctx context.Context, id int32) (Models.Subject, error)
	Insert(ctx context.Context, subject Models.Subject) (Models.Subject, error)
	Update(ctx context.Context, subject Models.Subject, id int32) (Models.Subject, error)
	Delete(ctx context.Context, id int32) error
}