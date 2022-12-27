package http

import (
	"github.com/MagonxESP/dropper/http/controller"
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
)

func RegisterRoutes(engine *gin.Engine) {
	auth := engine.Group("/oauth2")
	{
		auth.GET("/token", ginserver.HandleTokenRequest)
	}

	api := engine.Group("/api")
	{
		api.Use(ginserver.HandleTokenVerify())
		api.POST("/drop", controller.DropController)
		api.GET("/bucket/all", controller.GetAllBuckets)
	}
}
