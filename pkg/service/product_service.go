package service

import (
	"products/errs"
	"products/models"
	"products/pkg/repository"
)

func CreateProduct(req models.Product) error {
	err := repository.CreateProduct(req)
	if err != nil {
		return err
	}
	return nil
}

// func UpdateProduct(req models.Product, id uint) error {
// 	err := repository.UpdateProduct(req.ModelName, req.Description, req.Price, req.CategoryID, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func DeleteProductByID(productID uint) error {
	err := repository.DeleteProductByID(productID)
	if err != nil {
		if err == errs.ErrProductNotFound {
			return errs.ErrProductNotFound
		}
		return err
	}
	return nil
}

func GetProductByID(productID uint) (product []models.Product, err error) {
	product, err = repository.GetProductByID(productID)
	if err != nil {
		if err == errs.ErrProductNotFound {
			return nil, errs.ErrProductNotFound
		}
		return product, err
	}
	return product, nil
}

func GetProducts(minPrice, maxPrice int, description, categoryID int) ([]models.Product, error) {
	products, err := repository.GetAllProducts(minPrice, maxPrice, description, categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

