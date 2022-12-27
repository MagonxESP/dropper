package controller

import "github.com/gin-gonic/gin"

func ErrorJsonResponse(status int, err error, context *gin.Context) {
	context.JSON(status, map[string]string{
		"status": "error",
		"error":  err.Error(),
	})
}

func SuccessStatusJsonResponse(context *gin.Context) {
	context.JSON(200, map[string]string{
		"status": "success",
	})
}
