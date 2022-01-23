package controllers

import (
	"context"
	"net/http"
	"quiz-3/models"
	"quiz-3/services"
	"quiz-3/utils"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func IndexCategories(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var categories []models.Category

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	categories, err := services.GetCategories(ctx)

	if err != nil {
		var response map[string]string = map[string]string{
			"message": err.Error(),
		}

		utils.ResponseJSON(w, response, http.StatusInternalServerError)
		return
	}

	var response map[string]interface{} = map[string]interface{}{
		"message": "OK",
		"data":    categories,
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}

func ShowCategory(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	var category models.Category

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	id, _ := strconv.Atoi(params.ByName("id"))

	category, err := services.GetCategory(ctx, id)

	if err != nil {
		var response map[string]string = map[string]string{
			"message": err.Error(),
		}

		utils.ResponseJSON(w, response, http.StatusInternalServerError)
		return
	}

	var response map[string]interface{} = map[string]interface{}{
		"message": "OK",
		"data":    category,
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}

func StoreCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var category models.Category

	err := utils.ParseBody(r, &category)

	if err != nil {
		var response map[string]string = map[string]string{
			"message": err.Error(),
		}

		utils.ResponseJSON(w, response, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	_, err = services.StoreCategory(ctx, category)

	if err != nil {
		var response map[string]string = map[string]string{
			"message": err.Error(),
		}

		utils.ResponseJSON(w, response, http.StatusInternalServerError)
		return
	}

	var response map[string]interface{} = map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var category models.Category

	err := utils.ParseBody(r, &category)

	if err != nil {
		var response map[string]string = map[string]string{
			"message": err.Error(),
		}

		utils.ResponseJSON(w, response, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	id, _ := strconv.Atoi(params.ByName("id"))

	_, err = services.UpdateCategory(ctx, category, id)

	if err != nil {
		var response map[string]string = map[string]string{
			"message": err.Error(),
		}

		utils.ResponseJSON(w, response, http.StatusInternalServerError)
		return
	}

	var response map[string]interface{} = map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}

func DestroyCategory(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	id, _ := strconv.Atoi(params.ByName("id"))

	err := services.DeleteCategory(ctx, id)

	if err != nil {
		var response map[string]string = map[string]string{
			"message": err.Error(),
		}

		utils.ResponseJSON(w, response, http.StatusInternalServerError)
		return
	}

	var response map[string]interface{} = map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}
