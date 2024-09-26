// service/wishlist.go
package service

import (
	"products/models"
	"products/pkg/repository"
)

func AddToWishlist(wish models.Wishlist) error {
	if err := repository.CreateWishlist(wish); err != nil {
		return err
	}
	return nil
}

func RemoveFromWishlist(productID uint) error {
	if err := repository.RemoveItemFromWishlist(productID); err != nil {
		return err
	}

	return nil
}

func GetUserWishlistByID(id uint) ([]models.Wishlist, error) {
	var wishlists []models.Wishlist
	wishlists, err := repository.GetWishlistByID(id)
	if err != nil {
		return wishlists, err
	}
	return wishlists, nil
}
