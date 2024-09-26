package repository

import (
	"products/db"
	"products/errs"
	"products/models"

	"gorm.io/gorm"
)

func CreateProduct(product models.Product) error {
	if err := db.GetDBConn().Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(productID uint) (products []models.Product, err error) {

	if err := db.GetDBConn().First(products, productID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrProductNotFound
		}
		return nil, err
	}
	return products, nil
}

func GetAllProducts(minPrice, maxPrice int, description int, categoryID int) (products []models.Product, err error) {
	query := db.GetDBConn().Model(&models.Product{})

	if minPrice > 0 {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice > 0 {
		query = query.Where("price <= ?", maxPrice)
	}
	if description != 0 {
		query = query.Where("description = ?", description)
	}
	if categoryID != 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}



func DeleteProductByID(productID uint) error {
	if err := db.GetDBConn().Delete(&models.Product{}, productID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.ErrProductNotFound
		}
		return err
	}
	return nil
}
