package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tucazorron/confio-api/model"
)

type CountryController struct{}

func NewCountryController() CountryController {
	return CountryController{}
}

func (c *CountryController) GetCountries(ctx *gin.Context) {
	countries := []model.Country{
		{Code: "BRA", Name: "Brasil"},
		{Code: "ARG", Name: "Argentina"},
		{Code: "CHL", Name: "Chile"},
		{Code: "PER", Name: "Peru"},
	}
	ctx.JSON(http.StatusOK, countries)
}
