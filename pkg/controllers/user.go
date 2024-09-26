package controllers

import (
	"products/logger"
	"products/models"
	"products/errs"
	"products/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// PrintUsers
// @Summary Get All Users
// @Security ApiKeyAuth
// @Tags users
// @Description get list of all users
// @ID get-all-users
// @Produce json
// @Param q query string false "fill if you need search"
// @Param role query string true "fill if you need search"
// @Success 200 {array} models.User
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/users [get]
func GetAllUsers(c *gin.Context) {

	urole := c.GetString(userRoleCtx)
	if urole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	if urole != "admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}
	
	logger.Info.Printf("Client with ip: [%s] requested list of users\n", c.ClientIP())
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
	logger.Info.Printf("Client with ip: [%s] got list of users\n", c.ClientIP())
}

// GetUserByID
// @Summary Get User By ID
// @Security ApiKeyAuth
// @Tags users
// @Description get user by id
// @ID get-user-by-id
// @Accept json
// @Produce json
// @Param id path integer true "id of the user"
// @Param input body models.User true "user info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/users/{id} [get]
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] invalid user_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	user, err := service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)

}


// CreateUser
// @Summary Create User
// @Security ApiKeyAuth
// @Tags users
// @Description create new user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body models.User true "new route info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})

}

// UpdateUserByID
// @Summary Update User
// @Security ApiKeyAuth
// @Tags users
// @Description update existed user
// @ID update-user
// @Accept json
// @Produce json
// @Param id path integer true "id of the user"
// @Param input body models.User true "user update info"
// @Success 200 {object} defaultResponse
// @Failure 400 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/user/{id} [put]
func UpdateUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user.ID = int(id)

	c.JSON(http.StatusOK, user)
}
