package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"tugas-15/models"
)

type GradeRepositoryImpl struct {
	DB *sql.DB
}

func NewGradeRepository(db *sql.DB) GradeRepository {
	return &GradeRepositoryImpl{DB: db}
}

func (repository *GradeRepositoryImpl) GetAll(ctx context.Context) ([]models.Grade, error) {
	var grades []models.Grade = []models.Grade{}
	script := "SELECT * FROM grades"

	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return grades, err
	}

	for rows.Next() {
		var grade models.Grade
		err := rows.Scan(&grade.ID, &grade.StudentId, &grade.SubjectId, &grade.Score, &grade.Index, &grade.CreatedAt, &grade.UpdatedAt)

		if err != nil {
			return grades, errors.New(err.Error())
		}

		grades = append(grades, grade)
	}

	defer rows.Close()

	return grades, nil
}

func (repository *GradeRepositoryImpl) Get(ctx context.Context, id int32) (models.Grade, error) {
	var grade models.Grade
	script := "SELECT * FROM grades WHERE id = ?"

	err := repository.DB.QueryRowContext(ctx, script, id).Scan(&grade.ID, &grade.StudentId, &grade.SubjectId, &grade.Score, &grade.Index, &grade.CreatedAt, &grade.UpdatedAt)

	if err != nil {
		return grade, errors.New(err.Error())
	}

	return grade, nil
}

func (repository *GradeRepositoryImpl) Insert(ctx context.Context, grade models.Grade) (models.Grade, error) {
	script := "INSERT INTO grades (score, `index`, student_id, subject_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := repository.DB.PrepareContext(ctx, script)

	if err != nil {
		return grade, err
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	res, err := stmt.ExecContext(ctx, grade.Score, grade.Index, grade.StudentId, grade.SubjectId, now, now)

	if err != nil {
		return grade, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return grade, err
	}

	grade.ID = uint(id)

	defer stmt.Close()

	return grade, nil
}

func (repository *GradeRepositoryImpl) Update(ctx context.Context, grade models.Grade, id int32) (models.Grade, error) {
	script := "UPDATE grades SET score = ?, `index` = ?, student_id = ?, subject_id = ?, updated_at = ? WHERE id = ?"

	stmt, err := repository.DB.PrepareContext(ctx, script)

	if err != nil {
		return grade, err
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	_, err = stmt.ExecContext(ctx, grade.Score, grade.Index, grade.StudentId, grade.SubjectId, now, id)

	if err != nil {
		return grade, err
	}

	defer stmt.Close()

	return grade, nil
}

func (repository *GradeRepositoryImpl) Delete(ctx context.Context, id int32) error {
	script := "DELETE FROM grades WHERE id = ?"

	stmt, err := repository.DB.PrepareContext(ctx, script)

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, id)

	if err != nil {
		return err
	}

	defer stmt.Close()

	return nil
}
