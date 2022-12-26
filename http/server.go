package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

func StartHttpServer() {
	engine := gin.Default()
	RegisterRoutes(engine)
	err := engine.Run()

	if err != nil {
		log.Fatal(err)
	}
}
