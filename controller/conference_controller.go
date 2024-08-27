package controller

import (
	"net/http"
	"tucazorron/confio-api/model"
	"tucazorron/confio-api/usecase"

	"github.com/gin-gonic/gin"
)

type ConferenceController struct {
	conferenceUseCase usecase.ConferenceUseCase
}

func NewConferenceController(usecase usecase.ConferenceUseCase) ConferenceController {
	return ConferenceController{
		conferenceUseCase: usecase,
	}
}

func (ctrl *ConferenceController) GetConferences(ctx *gin.Context) {
	conferences, err := ctrl.conferenceUseCase.GetConferences()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, conferences)
}

func (ctrl *ConferenceController) CreateConference(ctx *gin.Context) {
	var conference model.Conference
	err := ctx.BindJSON(&conference)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	insertedConference, err := ctrl.conferenceUseCase.CreateConference(conference)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, insertedConference)
}
