package services

import (
	"context"
	"errors"
	"net/url"
	conf "quiz-3/config"
	"quiz-3/models"
	"quiz-3/repositories"
)

func GetBooks(ctx context.Context, query url.Values) ([]models.Book, error) {
	var books []models.Book

	repository := repositories.NewBookRepository(conf.DB)

	books, err := repository.GetBooks(ctx, query)

	if err != nil {
		return books, err
	}

	return books, nil
}

func GetBook(ctx context.Context, id int) (models.Book, error) {
	var book models.Book

	repository := repositories.NewBookRepository(conf.DB)

	book, err := repository.GetBookById(ctx, id)

	if err != nil {
		return book, err
	}

	return book, nil
}

func StoreBook(ctx context.Context, book models.Book) (models.Book, []error) {
	var newBook models.Book
	var errs []error

	repository := repositories.NewBookRepository(conf.DB)

	var errorMsgs map[string]string = book.ValidateBook()

	if len(errorMsgs) != 0 {
		for _, v := range errorMsgs {
			errs = append(errs, errors.New(v))
		}

		return newBook, errs
	}

	book.ConvertBook()

	newBook, err := repository.CreateBook(ctx, book)

	if err != nil {
		return newBook, errs
	}

	return newBook, nil
}

func UpdateBook(ctx context.Context, body models.Book, id int) (models.Book, []error) {
	var book models.Book
	var errs []error

	repository := repositories.NewBookRepository(conf.DB)

	book, err := GetBook(ctx, id)

	book.Title = body.Title
	book.Description = body.Description
	book.ImageUrl = body.ImageUrl
	book.ReleaseYear = body.ReleaseYear
	book.Price = body.Price
	book.TotalPage = body.TotalPage
	book.CategoryId = body.CategoryId

	if err != nil {
		errs = append(errs, err)
		return book, errs
	}

	var errorMsgs map[string]string = book.ValidateBook()

	if len(errorMsgs) != 0 {
		for _, v := range errorMsgs {
			errs = append(errs, errors.New(v))
		}

		return book, errs
	}

	book.ConvertBook()

	book, err = repository.UpdateBook(ctx, book)

	if err != nil {
		errs = append(errs, err)
		return book, errs
	}

	return book, nil
}

func DeleteBook(ctx context.Context, id int) error {
	repository := repositories.NewBookRepository(conf.DB)

	err := repository.DeleteBook(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
