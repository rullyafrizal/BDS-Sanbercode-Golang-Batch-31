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

func IndexBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var books []models.Book
	query := r.URL.Query()

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	books, err := services.GetBooks(ctx, query)

	if err != nil {
		var response map[string]interface{} = map[string]interface{}{
			"message": err.Error(),
		}

		utils.ResponseJSON(w, response, http.StatusInternalServerError)
		return
	}

	var response map[string]interface{} = map[string]interface{}{
		"message": "OK",
		"data":    books,
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}

func ShowBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	id, _ := strconv.Atoi(params.ByName("id"))

	book, err := services.GetBook(ctx, id)

	if err != nil {
		var response map[string]interface{} = map[string]interface{}{
			"message": err.Error(),
		}

		utils.ResponseJSON(w, response, http.StatusInternalServerError)
		return
	}

	var response map[string]interface{} = map[string]interface{}{
		"message": "OK",
		"data":    book,
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}

func StoreBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var book models.Book

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := utils.ParseBody(r, &book)

	if err != nil {
		var response map[string]interface{} = map[string]interface{}{
			"message": err.Error(),
		}

		utils.ResponseJSON(w, response, http.StatusBadRequest)
		return
	}

	_, errs := services.StoreBook(ctx, book)

	if errs != nil || len(errs) != 0 {
		var response map[string]interface{} = map[string]interface{}{
			"message": "Error",
			"data":    utils.StringifyErrors(errs...),
		}

		utils.ResponseJSON(w, response, http.StatusUnprocessableEntity)
		return
	}

	var response map[string]interface{} = map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var book models.Book

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := utils.ParseBody(r, &book)

	if err != nil {
		var response map[string]interface{} = map[string]interface{}{
			"message": err.Error(),
		}

		utils.ResponseJSON(w, response, http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(params.ByName("id"))

	_, errs := services.UpdateBook(ctx, book, id)

	if errs != nil || len(errs) != 0 {
		var response map[string]interface{} = map[string]interface{}{
			"message": "Error",
			"data":    utils.StringifyErrors(errs...),
		}

		utils.ResponseJSON(w, response, http.StatusUnprocessableEntity)
		return
	}

	var response map[string]interface{} = map[string]interface{}{
		"message": "OK",
	}

	utils.ResponseJSON(w, response, http.StatusOK)
}

func DestroyBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	id, _ := strconv.Atoi(params.ByName("id"))

	err := services.DeleteBook(ctx, id)

	if err != nil {
		var response map[string]interface{} = map[string]interface{}{
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
