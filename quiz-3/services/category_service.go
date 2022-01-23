package services

import (
	"context"
	conf "quiz-3/config"
	"quiz-3/models"
	"quiz-3/repositories"
)

func GetCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category

	repository := repositories.NewCategoryRepository(conf.DB)

	categories, err := repository.GetCategories(ctx)

	if err != nil {
		return categories, err
	}

	return categories, nil
}

func GetCategory(ctx context.Context, id int) (models.Category, error) {
	var category models.Category

	repository := repositories.NewCategoryRepository(conf.DB)

	category, err := repository.GetCategoryById(ctx, id)

	if err != nil {
		return category, err
	}

	return category, nil
}

func StoreCategory(ctx context.Context, category models.Category) (models.Category, error) {
	var newCategory models.Category

	repository := repositories.NewCategoryRepository(conf.DB)

	newCategory, err := repository.CreateCategory(ctx, category)

	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func UpdateCategory(ctx context.Context, body models.Category, id int) (models.Category, error) {
	var category models.Category

	repository := repositories.NewCategoryRepository(conf.DB)

	category, err := GetCategory(ctx, id)

	if err != nil {
		return category, err
	}

	category.Name = body.Name

	updatedCategory, err := repository.UpdateCategory(ctx, category)

	if err != nil {
		return updatedCategory, err
	}

	return updatedCategory, nil
}

func DeleteCategory(ctx context.Context, id int) error {
	repository := repositories.NewCategoryRepository(conf.DB)

	err := repository.DeleteCategory(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
