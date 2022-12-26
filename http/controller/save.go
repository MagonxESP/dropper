package controller

import (
	"github.com/MagonxESP/dropper/internal/application"
	"github.com/MagonxESP/dropper/internal/infraestructure"
	"github.com/MagonxESP/dropper/internal/infraestructure/pixiv"
	"github.com/gin-gonic/gin"
)

func SaveController(context *gin.Context) {
	source := context.Query("source")

	saver := application.NewRemoteFileSaver(
		pixiv.NewPixivIllustrationDownloader(),
		infraestructure.NewFileSystemWriter("images"),
	)

	err := saver.Save(source)

	if err != nil {
		context.JSON(500, map[string]string{
			"status": "error",
			"error":  err.Error(),
		})
	} else {
		context.JSON(200, map[string]string{
			"status": "ok",
		})
	}
}
