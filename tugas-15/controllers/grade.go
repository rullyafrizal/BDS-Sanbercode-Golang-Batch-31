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

func GetGrades(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	gradeRepository := repositories.NewGradeRepository(config.DB)

	grades, err := gradeRepository.GetAll(ctx)

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, grades, http.StatusOK)
}

func GetGrade(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	gradeRepository := repositories.NewGradeRepository(config.DB)

	id, _ := strconv.Atoi(p.ByName("id"))

	grade, err := gradeRepository.Get(ctx, int32(id))

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, grade, http.StatusOK)
}

func StoreGrade(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Use content-type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	gradeRepository := repositories.NewGradeRepository(config.DB)

	var grade models.Grade = models.Grade{}

	err := json.NewDecoder(r.Body).Decode(&grade)

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := grade.Validate(); err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	grade.ConvertScore()

	grade, err = gradeRepository.Insert(ctx, grade)

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusCreated)
}

func UpdateGrade(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Use content-type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	gradeRepository := repositories.NewGradeRepository(config.DB)

	var grade models.Grade = models.Grade{}

	err := json.NewDecoder(r.Body).Decode(&grade)

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := grade.Validate(); err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	grade.ConvertScore()

	id, _ := strconv.Atoi(p.ByName("id"))

	grade, err = gradeRepository.Update(ctx, grade, int32(id))

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}

func DestroyGrade(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	gradeRepository := repositories.NewGradeRepository(config.DB)

	id, _ := strconv.Atoi(p.ByName("id"))

	err := gradeRepository.Delete(ctx, int32(id))

	if err != nil {
		utils.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}
