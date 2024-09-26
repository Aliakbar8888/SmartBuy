package controllers

import (
	"net/http"
	"products/errs"
	"products/models"
	"products/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetWishlistByUserID
// @Summary Get Wishlist By ID
// @Security ApiKeyAuth
// @Tags wishlists
// @Description get wishlist by id
// @ID get-wishlist-by-id
// @Accept json
// @Produce json
// @Param id path integer true "id of the wishlist"
// @Param input body models.Wishlist true "wishlist info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/wishlist/{id} [get]
func GetWishlistByUserID(c *gin.Context) {
	
	id,err:=strconv.Atoi(c.Param("id"))
	if err!= nil{
		handleError(c, err)
		return
	}
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	wishlist, err := service.GetUserWishlistByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, wishlist)
}

// AddToWishlist
// @Summary Create Wishlist
// @Security ApiKeyAuth
// @Tags wishlists
// @Description create new wishlist
// @ID create-wishlist
// @Accept json
// @Produce json
// @Param input body models.Wishlist true "new wishlist info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/wishlist [post]
func AddToWishlist(c *gin.Context) {
	userID:=c.GetUint(userIDCtx)
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var wishlistItem models.Wishlist
	if err := c.BindJSON(&wishlistItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	wishlistItem.UserID=userID
	err := service.AddToWishlist(wishlistItem)
	if err != nil {
		return
	}
	// c.JSON(http.StatusOK, wishlistItem)
	c.JSON(http.StatusCreated, gin.H{"Message": "Wishlist is successfuly created"})
}


// RemoveFromWishlist
// @Summary Remove From Wishlist
// @Security ApiKeyAuth
// @Tags wishlists
// @Description remove product from wishlist
// @ID delete-wishlist-by-ID
// @Param id path integer true "id of the product"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/wishlist/{id} [delete]
func RemoveFromWishlist(c *gin.Context) {

	urole := c.GetString(userRoleCtx)
	if urole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	wishlistID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handleError(c, err)
		return
	}

	err = service.RemoveFromWishlist(uint(wishlistID))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from wishlist"})
}
