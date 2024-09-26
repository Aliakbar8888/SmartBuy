package db

import "products/models"

func Migrate() error {
	err := dbConn.AutoMigrate(models.User{},
		models.Category{},
		models.Product{},
		models.Order{},
		models.OrderItem{},
		models.Wishlist{},
	    models.Admin{},)
	if err != nil {
		return err
	}
	return nil
}
