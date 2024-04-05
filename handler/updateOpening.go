package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vagnermaltauro/gopportunities/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	r := UpdateOpeningRequest{}

	if err := ctx.BindJSON(&r); err != nil {
		logger.Errorf("error context: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := r.Validate(); err != nil {
		logger.Errorf("error validating r: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if r.Role != "" {
		opening.Role = r.Role
	}

	if r.Company != "" {
		opening.Company = r.Company
	}

	if r.Location != "" {
		opening.Location = r.Location
	}

	if r.Salary > 0 {
		opening.Salary = r.Salary
	}

	if r.Location != "" {
		opening.Location = r.Location
	}

	if r.Remote != nil {
		opening.Remote = *r.Remote
	}

	if r.Link != "" {
		opening.Link = r.Link
	}

	if r.Salary > 0 {
		opening.Salary = r.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
