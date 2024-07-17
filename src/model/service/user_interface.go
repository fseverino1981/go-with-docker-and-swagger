package service

import (
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/model"
	"go-with-docker-and-swagger/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUserServices(string, model.UserDomainInterface) *rest_err.RestErr

	FindUserByIDServices(string) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailServices(string) (model.UserDomainInterface, *rest_err.RestErr)

	findUserByEmailAndPasswordServices(string, string) (model.UserDomainInterface, *rest_err.RestErr)

	DeleteUserServices(string) *rest_err.RestErr

	LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
