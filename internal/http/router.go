package http

import (
	"github.com/MagonxESP/dropper/internal/http/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	engine.GET("/", controller.HomeController)
	engine.POST("/save", controller.SaveController)
}
