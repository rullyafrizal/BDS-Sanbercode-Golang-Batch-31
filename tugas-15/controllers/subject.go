package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"tugas-15/config"
	"tugas-15/models"
	"tugas-15/repositories"
	"tugas-15/utils"

	"github.com/julienschmidt/httprouter"
)

func GetSubjects(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	subjectRepository := repositories.NewSubjectRepository(config.DB)

	subjects, err := subjectRepository.GetAll(ctx)

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, subjects, http.StatusOK)
}

func GetSubject(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	subjectRepository := repositories.NewSubjectRepository(config.DB)

	id, _ := strconv.Atoi(p.ByName("id"))

	subject, err := subjectRepository.Get(ctx, int32(id))

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, subject, http.StatusOK)
}

func StoreSubject(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Use content-type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var subjectRepository repositories.SubjectRepository = repositories.NewSubjectRepository(config.DB)

	var subject models.Subject

	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	subject, err := subjectRepository.Insert(ctx, subject)

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}

func UpdateSubject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Use content-type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var subjectRepository repositories.SubjectRepository = repositories.NewSubjectRepository(config.DB)

	var subject models.Subject

	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(p.ByName("id"))

	subject, err := subjectRepository.Update(ctx, subject, int32(id))

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w,response, http.StatusOK)
}

func DestroySubject(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	subjectRepository := repositories.NewSubjectRepository(config.DB)

	id, _ := strconv.Atoi(p.ByName("id"))

	err := subjectRepository.Delete(ctx, int32(id))

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}