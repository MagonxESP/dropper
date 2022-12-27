package controller

import (
	"errors"
	"github.com/MagonxESP/dropper/internal/application"
	"github.com/MagonxESP/dropper/internal/infraestructure/downloader"
	"github.com/MagonxESP/dropper/internal/infraestructure/repository"
	"github.com/MagonxESP/dropper/internal/infraestructure/writer"
	"github.com/gin-gonic/gin"
)

type DropRequest struct {
	Source     string `json:"source"`
	BucketName string `json:"bucket_name"`
}

func (d *DropRequest) Validate() error {
	if d.Source == "" {
		return errors.New("the source parameter is missing")
	}

	if d.BucketName == "" {
		return errors.New("the bucket_name parameter is missing")
	}

	return nil
}

func DropController(context *gin.Context) {
	request, err := GetRequestJsonBody[DropRequest](context)

	if err != nil {
		ErrorJsonResponse(500, err, context)
		return
	}

	if err := request.Validate(); err != nil {
		ErrorJsonResponse(400, err, context)
		return
	}

	finder := application.NewBucketFinder(repository.NewYamlBucketRepository())
	bucket, err := finder.FindByName(request.BucketName)

	if err != nil {
		ErrorJsonResponse(500, err, context)
		return
	}

	writerInstance, err := writer.GetDownloadedFileWriterFromBucket(bucket)

	if err != nil {
		ErrorJsonResponse(400, err, context)
		return
	}

	saver := application.NewRemoteFileSaver(
		downloader.NewHttpFileDownloader(),
		writerInstance,
	)

	err = saver.Save(request.Source)

	if err != nil {
		ErrorJsonResponse(500, err, context)
	} else {
		SuccessStatusJsonResponse(context)
	}
}
