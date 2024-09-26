package service

import (
	"errors"
	"products/errs"
	"products/logger"
	"products/models"
	"products/pkg/repository"

	"gorm.io/gorm"
)

func CreateCategory(req models.Category) error {
	err := repository.CreateCategory(req)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategories() ([]models.Category, error) {
	// Логируем начало выполнения функции
	logger.Info.Println("Fetching all categories...")
	// Получаем категории из репозитория
	categories, err := repository.GetAllCategories()
	if err != nil {
		// Логируем ошибку
		logger.Error.Printf("Error fetching categories: %v", err)
		// Возвращаем пустой срез и ошибку
		return nil, errors.New("unable to retrieve categories at the moment")
	}

	// Если необходимо провести дополнительную обработку данных, можно сделать это здесь.
	// Например, можно фильтровать или сортировать категории.

	// Логируем успешное завершение
	logger.Info.Println("Successfully fetched all categories.")

	// Возвращаем категории и nil как ошибку
	return categories, nil
}

func UpdateCategory(id uint,categories models.Category) error {
	err := repository.UpdateCategory(categories.ID, categories.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrCategoryNotFound
		}
		return err
	}
	return nil
}

func DeleteCategory(categoryID uint) error {
	err := repository.DeleteCategory(categoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrCategoryNotFound
		}
		return err
	}
	return nil
}

func GetCategoryByID(id uint) (categories []models.Category, err error) {
	categories, err = repository.GetCategoryByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrCategoryNotFound
		}
		return nil, err
	}

	return categories, nil
}
