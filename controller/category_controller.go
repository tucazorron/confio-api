package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tucazorron/confio-api/model"
)

type CategoryController struct{}

func NewCategoryController() CategoryController {
	return CategoryController{}
}

func (c *CategoryController) GetCategories(ctx *gin.Context) {
	categories := []model.Category{
		{ID: 1, Name: "Ciencia da Computacao"},
		{ID: 2, Name: "Arquitetura e Urbanismo"},
		{ID: 3, Name: "Engenharia Eletrica"},
		{ID: 4, Name: "Matematica"},
	}
	ctx.JSON(http.StatusOK, categories)
}
