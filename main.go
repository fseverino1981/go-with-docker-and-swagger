package main

import (
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/controller"
	"go-with-docker-and-swagger/src/controller/routes"
	"go-with-docker-and-swagger/src/model/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Init Services
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
