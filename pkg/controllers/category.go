package controllers

import (
	"strconv"
	"net/http"
	"products/errs"
	"products/models"
	"products/pkg/service"
	"github.com/gin-gonic/gin"
)

// CreateCategory
// @Summary Create Category
// @Security ApiKeyAuth
// @Tags categories
// @Description create new category
// @ID create-category
// @Accept json
// @Produce json
// @Param input body models.Category true "new category info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/category [post]
func CreateCategory(c *gin.Context) {
    
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}
	var req models.Category
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := service.CreateCategory(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

// GetAllCategories
// @Summary Get All Categories
// @Security ApiKeyAuth
// @Tags categories
// @Description get list of all categories
// @ID get-all-categpries
// @Produce json
// @Param q query string false "fill if you need search"
// @Success 200 {array} models.Category
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/category [get]
func GetAllCategories(c *gin.Context) {

    urole := c.GetString(userRoleCtx)
	if urole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	
    categories, err := service.GetAllCategories()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch categories"})
        return
    }

    c.JSON(http.StatusOK, categories)
}

// UpdateCategoryByID
// @Summary Update Category
// @Security ApiKeyAuth
// @Tags categories
// @Description update existed category
// @ID update-category
// @Accept json
// @Produce json
// @Param id path integer true "id of the category"
// @Param input body models.Category true "category update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/category/{id} [put]
func UpdateCategoryByID(c *gin.Context) {
    categoryIDStr := c.Param("id")
    var req models.Category

    urole := c.GetString(userRoleCtx)
	if urole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

    categoryID, err := strconv.ParseUint(categoryIDStr,10,32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

 
    err = service.UpdateCategory(uint(categoryID),req)
    if err != nil {
        if err == errs.ErrCategoryNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update category"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

// DeleteCategoryByID
// @Summary Delete Category By ID
// @Security ApiKeyAuth
// @Tags categories
// @Description delete category by ID
// @ID delete-category-by-ID
// @Param id path integer true "id of the category"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/category/{id} [delete]
func DeleteCategoryByID(c *gin.Context) {
    categoryIDStr := c.Param("id")

	urole := c.GetString(userRoleCtx)
	if urole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

    categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    err = service.DeleteCategory(uint(categoryID))
    if err != nil {
        if err == errs.ErrCategoryNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete category"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}


// GetCategoryByID
// @Summary Get Category By ID
// @Security ApiKeyAuth
// @Tags categories
// @Description get category by id
// @ID get-category-by-id
// @Accept json
// @Produce json
// @Param id path integer true "id of the category"
// @Param input body models.Category true "category info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/category/{id} [get]
func GetCategoryByID(c *gin.Context) {
    categoryIDStr := c.Param("id")
    
	urole := c.GetString(userRoleCtx)
	if urole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}
    
    categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    category, err := service.GetCategoryByID(uint(categoryID))
    if err != nil {
        if err == errs.ErrCategoryNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch category"})
        return
    }

    c.JSON(http.StatusOK, category)
}
