package controllers

import (
	"products/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}

type defaultResponse struct {
	Message string `json:"message"`
}

func newDefaultResponse(message string) defaultResponse {
	return defaultResponse{
		Message: message,
	}
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func handleError(c *gin.Context, err error) {
	if errors.Is(err, errs.ErrUsernameUniquenessFailed) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.ErrValidationFailed) {
		c.JSON(http.StatusBadRequest, newDefaultResponse(err.Error()))
	} else if errors.Is(err, errs.ErrPermissionDenied) {
		c.JSON(http.StatusForbidden, newDefaultResponse(err.Error()))
	} else if errors.Is(err, errs.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, newDefaultResponse(err.Error()))
	} else {
		c.JSON(http.StatusInternalServerError, newErrorResponse(errs.ErrSomethingWentWrong.Error()))
	}
 }