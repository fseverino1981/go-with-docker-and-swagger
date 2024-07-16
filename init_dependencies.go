package main

import (
	"go-with-docker-and-swagger/src/controller"
	"go-with-docker-and-swagger/src/model/repository"
	"go-with-docker-and-swagger/src/model/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	return userController
}
