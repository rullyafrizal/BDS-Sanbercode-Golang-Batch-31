package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
	Models "tugas-15/models"
)

type SubjectRepositoryImpl struct {
	DB *sql.DB
}

func NewSubjectRepository(db *sql.DB) SubjectRepository {
	return &SubjectRepositoryImpl{DB: db}
}

func (repository *SubjectRepositoryImpl) GetAll(ctx context.Context) ([]Models.Subject, error) {
	script := "SELECT * FROM subjects"
	var subjects []Models.Subject = []Models.Subject{}

	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		log.Fatal(err.Error())
		return subjects, errors.New(err.Error())
	}

	for rows.Next() {
		var subject Models.Subject

		err = rows.Scan(&subject.ID, &subject.Name, &subject.CreatedAt, &subject.UpdatedAt)

		if err != nil {
			log.Fatal(err.Error())
			return subjects, errors.New(err.Error())
		}

		subjects = append(subjects, subject)
	}

	defer rows.Close()

	return subjects, nil
}

func (repository *SubjectRepositoryImpl) Get(ctx context.Context, id int32) (Models.Subject, error) {
	var subject Models.Subject = Models.Subject{}
	script := "SELECT * FROM subjects WHERE id = ? LIMIT 1;"

	row, err := repository.DB.QueryContext(ctx, script, id)

	if err != nil {
		log.Fatal(err.Error())
		return subject, errors.New(err.Error())
	}

	if row.Next() {
		row.Scan(&subject.ID, &subject.Name, &subject.CreatedAt, &subject.UpdatedAt)
	}

	defer row.Close()

	return subject, nil
}

func (repository *SubjectRepositoryImpl) Insert(ctx context.Context, subject Models.Subject) (Models.Subject, error) {
	script := "INSERT INTO subjects (name, created_at, updated_at) VALUES (?, ?, ?)"

	stmt, err := repository.DB.PrepareContext(ctx, script)

	if err != nil {
		log.Fatal(err.Error())
		return subject, errors.New(err.Error())
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	result, err := stmt.ExecContext(ctx, subject.Name, now, now)

	if err != nil {
		log.Fatal(err.Error())
		return subject, errors.New(err.Error())
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err.Error())
		return subject, errors.New(err.Error())
	}

	subject.ID = uint(id)

	defer stmt.Close()

	return subject, nil
}

func (repository *SubjectRepositoryImpl) Update(ctx context.Context, subject Models.Subject, id int32) (Models.Subject, error) {
	script := "UPDATE subjects SET name = ?, updated_at = ? WHERE id = ?"

	now := time.Now().Format("2006-01-02 15:04:05")

	result, err := repository.DB.ExecContext(ctx, script, subject.Name, now, id)

	if err != nil {
		log.Fatal(err.Error())
		return subject, errors.New(err.Error())
	}

	affected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
		return subject, errors.New(err.Error())
	}

	if affected == 0 {
		return subject, errors.New("subject not found")
	}

	return subject, nil
}

func (repository *SubjectRepositoryImpl) Delete(ctx context.Context, id int32) error {
	script := "DELETE FROM subjects WHERE id = ?;"

	result, err := repository.DB.ExecContext(ctx, script, id)

	if err != nil {
		log.Fatal(err.Error())
		return errors.New(err.Error())
	}

	affected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
		return errors.New(err.Error())
	}

	if affected == 0 {
		return errors.New("subject not found")
	}

	return nil
}
