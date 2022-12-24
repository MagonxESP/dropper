package controller

import "github.com/gin-gonic/gin"

func HomeController(context *gin.Context) {
	context.JSON(200, map[string]string{
		"hello": "world!",
	})
}
