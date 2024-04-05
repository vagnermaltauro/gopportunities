package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vagnermaltauro/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Create a new opening
// @Description Create a new opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param body body CreateOpeningRequest true "CreateOpeningRequest"
// @Success 201 {object} CreateOpeningResponse
// @Error 400 {object} ErrorResponse
// @Error 500 {object} ErrorResponse
// @Router /openings [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("Error binding request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error create opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
