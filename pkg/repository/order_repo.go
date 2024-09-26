package repository

import (
	"products/db"
	"products/models"
)

func CreateOrder(order models.Order) error {
	if err := db.GetDBConn().Create(&order).Error; err != nil {
		return err
	}
	return nil
}


func GetOrdersByUserID(id uint) ([]models.Order, error) {
	var orders []models.Order
	if err := db.GetDBConn().Preload("Product").Joins("join products ON products.id=orders.product_id").Where("orders.id = ?", id).Find(&orders).Error; err != nil {
		return nil,
			err
	}
	return orders, nil
}

func UpdateOrder(productid uint, quant int, id, userid uint) error {
	if err := db.GetDBConn().Save(&models.Order{ID: id, UserID: userid, ProductID: productid, Quantity: quant}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOrder(orderID uint) error {
	if err := db.GetDBConn().Where("id = ?", orderID).Delete(&models.Order{}).Error; err != nil {
		return err
	}
	return nil
}

func GetUserOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := db.GetDBConn().Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetAllOrders() (orders []models.Order, err error) {

	if err := db.GetDBConn().Preload("Product").Joins("join products ON products.id=orders.product_id").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
