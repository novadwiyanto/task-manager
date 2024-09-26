package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorValidate(ctx *gin.Context, err error) {
	validationErrors := err.(validator.ValidationErrors)
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"message": validationErrors[0].Translate,
		"data":    nil,
	})
}

func ReturnError(err error) error {
	if err != nil {
		return err
	}

	return nil
}
