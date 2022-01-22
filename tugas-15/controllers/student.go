package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"tugas-15/config"
	"tugas-15/models"
	"tugas-15/repositories"
	"tugas-15/utils"

	"github.com/julienschmidt/httprouter"
)

func GetStudents(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var studentRepository repositories.StudentRepository = repositories.NewStudentRepository(config.DB)

	students, err := studentRepository.GetAll(ctx)

	if err != nil {
		log.Fatal("Error : ", err.Error())
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
	}

	utils.ResponseJSON(w, students, http.StatusOK)
}

func GetStudent(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var studentRepository repositories.StudentRepository = repositories.NewStudentRepository(config.DB)

	id, _ := strconv.Atoi(p.ByName("id"))

	student, err := studentRepository.Get(ctx, int32(id))

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, student, http.StatusOK)
}

func StoreStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Use content-type application/json", http.StatusBadRequest)
		return
	}

	var studentRepository repositories.StudentRepository = repositories.NewStudentRepository(config.DB)
	var student models.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	_, err := studentRepository.Insert(ctx, student)

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusCreated)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Use content-type application/json", http.StatusBadRequest)
		return
	}

	var studentRepository repositories.StudentRepository = repositories.NewStudentRepository(config.DB)
	var student models.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(p.ByName("id"))

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	_, err := studentRepository.Update(ctx, student, int32(id))

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}

func DestroyStudent(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	var studentRepository repositories.StudentRepository = repositories.NewStudentRepository(config.DB)

	id, _ := strconv.Atoi(p.ByName("id"))

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := studentRepository.Delete(ctx, int32(id))

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}
