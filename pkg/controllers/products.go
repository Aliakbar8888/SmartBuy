package controllers

import (
	"net/http"
	"products/errs"
	"products/models"
	"products/pkg/service"
	"strconv"
	"products/logger"
	"github.com/gin-gonic/gin"
)

// CreateProduct
// @Summary Create Product
// @Security ApiKeyAuth
// @Tags products
// @Description create new product
// @ID create-product
// @Accept json
// @Produce json
// @Param input body models.Product true "new product info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/product [post]
func CreateProduct(c *gin.Context) {
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	var req models.Product
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := service.CreateProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Create product successful"})
}

// GetAllProducts
// @Summary Get All Products
// @Security ApiKeyAuth
// @Tags products
// @Description get list of all products
// @ID get-all-products
// @Produce json
// @Param q query string false "fill if you need search"
// @Param min-price query int true "fill if you need search"
// @Param max-price query int true "fill if you need search"
// @Param description query int true "fill if you need search"
// @Param category query int true "fill if you need search"

// @Success 200 {array} models.Product
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/product [get]

	func GetAllProducts(c *gin.Context) {

		urole := c.GetString(userRoleCtx)
		if urole == "" {
			handleError(c, errs.ErrValidationFailed)
			return
		}
		minPriceStr := c.Query("min-price")
		maxPriceStr := c.Query("max-price")
		description := c.Query("description")
		categoryID := c.Query("category")
	
		logger.Info.Printf("[controllers.GetAllProducts] Request to get products with minPrice: %s, maxPrice: %s, description: %s, categoryID: %s\n", minPriceStr, maxPriceStr, description, categoryID)
	
		var minPrice, maxPrice int
		
			minPrice, err := strconv.Atoi(minPriceStr)
			if err != nil {
				logger.Error.Printf("[controllers.GetAllProducts] Error converting minPrice to int: %s", err.Error())
				handleError(c, errs.ErrIDIsNotCorrect)
				return
			
		}
	
			maxPrice, err = strconv.Atoi(maxPriceStr)
			if err != nil {
				logger.Error.Printf("[controllers.GetAllProducts] Error converting maxPrice to int: %s", err.Error())
				handleError(c, errs.ErrIncorrectInput)
				return
			
		}
			
			var categoryp int
			if categoryID != "" {
				categoryp, err = strconv.Atoi(categoryID)
				if err != nil {
					logger.Error.Printf("[controllers.GetAllProducts] Error converting categoryID to int: %s", err.Error())
					handleError(c, errs.ErrIncorrectInput)
					return
				}
			}

			var descriptionp int
			if description != "" {
				descriptionp, err = strconv.Atoi(description)
				if err != nil {
					logger.Error.Printf("[controllers.GetAllProducts] Error converting description to int: %s", err.Error())
					handleError(c, errs.ErrIncorrectInput)
					return
				}
			}


		products, err := service.GetProducts(minPrice, maxPrice, descriptionp, categoryp)
		if err != nil {
			handleError(c, err)
			return
		}
	
		logger.Info.Printf("[controllers.GetAllProducts] Successfully retrieved products: %d, maxPrice: %d, description: %s, categoryID: %s\n", minPrice, maxPrice, description, categoryID)
		c.JSON(http.StatusOK, products)
	}
 


// GetProductByID
// @Summary Get Product By ID
// @Security ApiKeyAuth
// @Tags products
// @Description get product by id
// @ID get-product-by-id
// @Accept json
// @Produce json
// @Param id path integer true "id of the product"
// @Param input body models.Product true "product info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/product/{id} [get]
func GetProductByID(c *gin.Context) {

	urole := c.GetString(userRoleCtx)
	if urole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	productID := c.Param("id")
	productIdUint, err := strconv.Atoi(productID)
	if err != nil {
		return
	}
	product, err := service.GetProductByID(uint(productIdUint))
	if err != nil {
		if err == errs.ErrProductNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch product"})
		return
	}
	c.JSON(http.StatusOK, product)
} 

// DeleteProductByID
// @Summary Delete Product By ID
// @Security ApiKeyAuth
// @Tags products
// @Description delete product by ID
// @ID delete-product-by-ID
// @Param id path integer true "id of the product"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/product/{id} [delete]
func DeleteProductByID(c *gin.Context) {

	urole := c.GetString(userRoleCtx)
	if urole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	productID := c.Param("id")
	productIdUint, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = service.DeleteProductByID(uint(productIdUint))
	if err != nil {
		if err == errs.ErrProductNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product successfully deleted"})
}
