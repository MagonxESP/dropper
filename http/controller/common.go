package controller

import (
	json2 "encoding/json"
	"github.com/gin-gonic/gin"
	"io"
)

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

func SuccessJsonResponse(data interface{}, context *gin.Context) {
	json := map[string]interface{}{
		"data": data,
	}

	context.JSON(200, json)
}

func GetRequestJsonBody[RequestStruct interface{}](context *gin.Context) (*RequestStruct, error) {
	json, err := io.ReadAll(context.Request.Body)

	if err != nil {
		return nil, err
	}

	var decoded RequestStruct
	err = json2.Unmarshal(json, &decoded)

	if err != nil {
		return nil, err
	}

	return &decoded, nil
}
