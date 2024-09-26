package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func JsonResponse(ctx *gin.Context, code int, msg string, data interface{}) {
	webResponse := response{
		Code:   code,
		Status: msg,
		Data:   data,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
