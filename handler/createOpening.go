package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("Error binding request: %v", err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("Error create opening: %v", err.Error())
		return
	}
}
