package repository

import (
	"errors"
	"products/db"
	"products/errs"
	"products/models"

	"gorm.io/gorm"
)

func CreateCategory(category models.Category) error {
	if err := db.GetDBConn().Create(category).Error; err != nil {
		return err
	}
	return nil
}

func GetCategoryByID(categoryID uint) (categories []models.Category, err error) {
	if err := db.GetDBConn().First(&categories, categoryID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return categories, errs.ErrCategoryNotFound
		}
		return categories, err
	}
	return categories, nil
}

func GetAllCategories() (categories []models.Category, err error) {
	if err := db.GetDBConn().Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func UpdateCategory(id uint, name string) error {
	err := db.GetDBConn().Save(&models.Category{ID: id, Name: name}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(categoryID uint) error {
	if err := db.GetDBConn().Delete(&models.Category{}, categoryID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrCategoryNotFound
		}
		return err
	}
	return nil
}
