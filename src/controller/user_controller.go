package controller

import (
	"go-with-docker-and-swagger/src/model/service"

	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{service: serviceInterface}
}

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
