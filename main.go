package main

import (
	"context"
	"go-with-docker-and-swagger/src/configuration/database/mongodb"
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/controller"
	"go-with-docker-and-swagger/src/controller/routes"
	"go-with-docker-and-swagger/src/model/repository"
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

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}

	//Init Services
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
