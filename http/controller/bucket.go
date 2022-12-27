package controller

import (
	"github.com/MagonxESP/dropper/internal/application"
	"github.com/MagonxESP/dropper/internal/infraestructure/repository"
	"github.com/gin-gonic/gin"
)

func GetAllBuckets(context *gin.Context) {
	finder := application.NewBucketFinder(repository.NewYamlBucketRepository())
	buckets, err := finder.All()

	if err != nil {
		ErrorJsonResponse(500, err, context)
		return
	}

	SuccessJsonResponse(buckets, context)
}
