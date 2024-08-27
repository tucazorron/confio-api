package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tucazorron/confio-api/model"
)

type CityController struct{}

func NewCityController() CityController {
	return CityController{}
}

func (c *CityController) GetCities(ctx *gin.Context) {
	cities := []model.City{
		{ID: 1, Name: "Brasilia", CountryCode: "BRA"},
		{ID: 2, Name: "Ouro Preto", CountryCode: "BRA"},
		{ID: 3, Name: "Cordoba", CountryCode: "ARG"},
		{ID: 4, Name: "Santiago", CountryCode: "CHL"},
	}
	ctx.JSON(http.StatusOK, cities)
}
