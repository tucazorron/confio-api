package main

import (
	"github.com/gin-gonic/gin"
	"tucazorron/confio-api/controller"
	"tucazorron/confio-api/db"
	"tucazorron/confio-api/repository"
	"tucazorron/confio-api/usecase"
)

func main() {
	router := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	ConferenceRepository := repository.NewProductRepository(dbConnection)
	ConferenceUseCase := usecase.NewConferenceUsecase(ConferenceRepository)
	ConferenceController := controller.NewConferenceController(ConferenceUseCase)
	router.GET("/conferences", ConferenceController.GetConferences)
	router.POST("/conference", ConferenceController.CreateConference)
	router.Run(":8000")
}
