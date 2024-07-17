package repository

import (
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(userID string, userDomain model.UserDomainInterface) *rest_err.RestErr

	DeleteUser(userID string) *rest_err.RestErr
}
