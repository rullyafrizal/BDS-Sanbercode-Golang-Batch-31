package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"tugas-14/config"
	"tugas-14/models"
	"tugas-14/repositories"
	"tugas-14/requests"
	"tugas-14/utils"

	"github.com/julienschmidt/httprouter"
)

func GetStudents(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var studentRepository *repositories.StudentRepositoryImpl = repositories.NewStudentRepository(config.DB)

	students, err := studentRepository.GetAll(ctx)

	if err != nil {
		log.Fatal("Error : ", err.Error())
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
	}

	utils.ResponseJSON(w, students, http.StatusOK)
}

func StoreStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Use content-type application/json", http.StatusBadRequest)
		return
	}

	var studentRepository *repositories.StudentRepositoryImpl = repositories.NewStudentRepository(config.DB)
	var student models.Student
	var request requests.StoreStudentRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := validate(request); err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	student = models.Student{
		Name:    request.Name,
		Grade:   request.Grade,
		Subject: request.Subject,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	determineIndex(&student)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	data, err := studentRepository.Insert(ctx, student)

	if err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "OK",
		"data":    data,
	}

	utils.ResponseJSON(w, response, http.StatusCreated)
}

func validate(request requests.StoreStudentRequest) error {
	if request.Grade > 100 {
		return errors.New("grade should no more than 100")
	}
	return nil
}

func determineIndex(student *models.Student) {
	if student.Grade >= 80 {
		student.Index = "A"
	} else if student.Grade >= 70 && student.Grade < 80 {
		student.Index = "B"
	} else if student.Grade >= 60 && student.Grade < 70 {
		student.Index = "C"
	} else if student.Grade >= 50 && student.Grade < 60 {
		student.Index = "D"
	} else if student.Grade < 50 {
		student.Index = "E"
	}
}
