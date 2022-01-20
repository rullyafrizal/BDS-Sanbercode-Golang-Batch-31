package repositories

import (
	"context"
	"database/sql"
	Models "tugas-14/models"
)

type StudentRepositoryImpl struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepositoryImpl {
	return &StudentRepositoryImpl{DB: db}
}

func (repository *StudentRepositoryImpl) Insert(ctx context.Context, student Models.Student) (Models.Student, error) {
	script := "INSERT INTO students (name, subject, grade, `index`) VALUES (?, ?, ?, ?)"
	stmt, err := repository.DB.PrepareContext(ctx, script)

	if err != nil {
		return student, err
	}

	result, err := stmt.ExecContext(ctx, student.Name, student.Subject, student.Grade, student.Index)

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

		err = rows.Scan(&student.ID, &student.Name, &student.Subject, &student.Grade, &student.Index)

		if err != nil {
			panic(err)
		}

		students = append(students, student)
	}

	defer rows.Close()

	return students, nil
}
