package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
	Models "tugas-15/models"
)

type StudentRepositoryImpl struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) StudentRepository {
	return &StudentRepositoryImpl{DB: db}
}

func (repository *StudentRepositoryImpl) Insert(ctx context.Context, student Models.Student) (Models.Student, error) {
	script := "INSERT INTO students (name, created_at, updated_at) VALUES (?, ?, ?)"
	stmt, err := repository.DB.PrepareContext(ctx, script)

	if err != nil {
		return student, err
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	result, err := stmt.ExecContext(ctx, student.Name, now, now)

	if err != nil {
		return student, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return student, err
	}

	student.ID = uint(id)

	defer stmt.Close()

	return student, nil
}

func (repository *StudentRepositoryImpl) GetAll(ctx context.Context) ([]Models.Student, error) {
	var students []Models.Student = []Models.Student{}
	script := "SELECT * FROM students"

	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return students, err
	}

	for rows.Next() {
		var student Models.Student

		err = rows.Scan(&student.ID, &student.Name, &student.CreatedAt, &student.UpdatedAt)

		if err != nil {
			panic(err)
		}

		students = append(students, student)
	}

	defer rows.Close()

	return students, nil
}

func (repository *StudentRepositoryImpl) Get(ctx context.Context, id int32) (Models.Student, error) {
	var student Models.Student = Models.Student{}
	script := "SELECT * FROM students WHERE id = ? LIMIT 1;"

	row, err := repository.DB.QueryContext(ctx, script, id)

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	defer row.Close()

	if row.Next() {
		row.Scan(&student.ID, &student.Name, &student.CreatedAt, &student.UpdatedAt)

		return student, nil
	}

	return student, errors.New("student not found")
}

func (repository *StudentRepositoryImpl) Update(ctx context.Context, student Models.Student, id int32) (Models.Student, error) {
	script := "UPDATE students SET name = ?, updated_at = ? WHERE id = ?;"

	now := time.Now().Format("2006-01-02 15:04:05")

	result, err := repository.DB.ExecContext(ctx, script, student.Name, now, id)

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	if affected == 0 {
		return student, errors.New("student not found")
	}

	return student, nil
}

func (repository *StudentRepositoryImpl) Delete(ctx context.Context, id int32) error {
	script := "DELETE FROM students WHERE id = ?;"

	result, err := repository.DB.ExecContext(ctx, script, id)

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	if affected == 0 {
		return errors.New("student not found")
	}

	return nil
}
