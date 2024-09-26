package db

import (
	"products/models"
)

func InsertSeeds() error {
	var productCategories []models.Category
	productCategories = append(productCategories, models.Category{Name: "Samsung"})
	productCategories = append(productCategories, models.Category{Name: "iPhone"})
	productCategories = append(productCategories, models.Category{Name: "Honor"})
	productCategories = append(productCategories, models.Category{Name: "Huawei"})
	productCategories = append(productCategories, models.Category{Name: "Redmi"})
	if err := dbConn.Create(&productCategories).Error; err != nil {
		return err
	}

	var model []models.Product
	model = append(model, models.Product{ModelName: "Samsung Zfold6", Description: 2024, Price: 9200, CategoryID: 1})
	model = append(model, models.Product{ModelName: "Samsung Zflip6", Description: 2024, Price: 8700, CategoryID: 1})
	model = append(model, models.Product{ModelName: "Samsung A35", Description: 2024, Price: 2900, CategoryID: 1})
	model = append(model, models.Product{ModelName: "Samsung A55", Description: 2024, Price: 3600, CategoryID: 1})
	model = append(model, models.Product{ModelName: "Samsung S24", Description: 2024, Price: 9500, CategoryID: 1})

	model = append(model, models.Product{ModelName: "iPhone 12", Description: 2020, Price: 4000,CategoryID: 2})
	model = append(model, models.Product{ModelName: "iPhone 13", Description: 2021, Price: 6000, CategoryID: 2})
	model = append(model, models.Product{ModelName: "iPhone 14", Description: 2022, Price: 7500, CategoryID: 2})
	model = append(model, models.Product{ModelName: "iPhone 15", Description: 2023, Price: 9000, CategoryID: 2})
	model = append(model, models.Product{ModelName: "iPhone 16", Description: 2024, Price: 10000, CategoryID: 2})

	model = append(model, models.Product{ModelName: "Honor 90", Description: 2021, Price: 2400,CategoryID: 3})
	model = append(model, models.Product{ModelName: "Honor X9b", Description: 2022, Price: 3600, CategoryID: 3})
	model = append(model, models.Product{ModelName: "Honor 200", Description: 2024, Price: 4100, CategoryID: 3})
	model = append(model, models.Product{ModelName: "Honor 200 Lite", Description: 2024, Price: 4800, CategoryID: 3})
	model = append(model, models.Product{ModelName: "Honor 200 Pro", Description: 2024, Price: 4400, CategoryID: 3})

	model = append(model, models.Product{ModelName: "Huawei 40", Description: 2020, Price: 1200,CategoryID: 4})
	model = append(model, models.Product{ModelName: "Huawei 90", Description: 2021, Price: 1600, CategoryID: 4})
	model = append(model, models.Product{ModelName: "Huawei Nova 11", Description: 2024, Price: 5500, CategoryID: 4})
	model = append(model, models.Product{ModelName: "Huawei Nova 9", Description: 2022, Price: 3000, CategoryID: 4})
	model = append(model, models.Product{ModelName: "Huawei Nova 10", Description: 2023, Price: 4000, CategoryID: 4})

	model = append(model, models.Product{ModelName: "Redmi Note 13", Description: 2024, Price: 4500,CategoryID: 5})
	model = append(model, models.Product{ModelName: "Redmi Note 10", Description: 2021, Price: 1300, CategoryID: 5})
	model = append(model, models.Product{ModelName: "Redmi Poco X6", Description: 2024, Price: 3500, CategoryID: 5})
	model = append(model, models.Product{ModelName: "Redmi Poco F5", Description: 2023, Price: 4000, CategoryID: 5})
	model = append(model, models.Product{ModelName: "Redmi Poco X4", Description: 2022, Price: 5000, CategoryID: 5})

	if err := dbConn.Create(&model).Error; err != nil {
		return err
	}
	return nil

}
