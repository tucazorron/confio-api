package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tucazorron/confio-api/model"
)

type OrganizerController struct{}

func NewOrganizerController() OrganizerController {
	return OrganizerController{}
}

func (c *OrganizerController) GetOrganizers(ctx *gin.Context) {
	organizers := []model.Organizer{
		{ID: 1, Name: "FAP", Email: "fap@email.com"},
		{ID: 2, Name: "CIC", Email: "cic@email.com"},
		{ID: 3, Name: "ABRUEM", Email: "abruem@email.com"},
		{ID: 4, Name: "SBPC", Email: "sbpc@email.com"},
	}
	ctx.JSON(http.StatusOK, organizers)
}
