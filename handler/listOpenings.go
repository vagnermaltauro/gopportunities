package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vagnermaltauro/gopportunities/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error Listing Openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
