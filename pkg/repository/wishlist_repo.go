package repository

import (
	"errors"
	"products/db"
	"products/errs"
	"products/models"

	"gorm.io/gorm"
)

func CreateWishlist(wishlist models.Wishlist) error {
	if err := db.GetDBConn().Create(&wishlist).Error; err != nil {
		return err
	}
	return nil
}

func GetWishlistByID(id uint) (wishlists[]models.Wishlist, err error) {
	if err := db.GetDBConn().Preload("Product").Joins("join products ON products.id=wishlists.product_id").Where("wishlists.id = ?", id).Find(&wishlists).Error; err != nil {
		return nil,
			err
	}
	return wishlists, nil
}


func RemoveItemFromWishlist(wishlistItemID uint) error {
	if err := db.GetDBConn().Where("id=?", wishlistItemID).Delete(&models.Wishlist{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrWishlistItemNotFound
		}
		return err
	}
	return nil
}
