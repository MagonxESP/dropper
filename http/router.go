package http

import (
	controller2 "github.com/MagonxESP/dropper/http/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	engine.GET("/", controller2.HomeController)
	engine.POST("/save", controller2.SaveController)
}
