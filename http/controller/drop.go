package controller

import (
	"errors"
	"github.com/MagonxESP/dropper/internal/application"
	"github.com/MagonxESP/dropper/internal/infraestructure"
	"github.com/MagonxESP/dropper/internal/infraestructure/downloader"
	"github.com/gin-gonic/gin"
)

func DropController(context *gin.Context) {
	source := context.Query("source")

	if source == "" {
		ErrorJsonResponse(400, errors.New("the source parameter is missing"), context)
		return
	}

	saver := application.NewRemoteFileSaver(
		downloader.NewHttpFileDownloader(),
		infraestructure.NewFileSystemWriter("images"),
	)

	err := saver.Save(source)

	if err != nil {
		ErrorJsonResponse(500, err, context)
	} else {
		SuccessStatusJsonResponse(context)
	}
}
