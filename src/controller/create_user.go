package controller

import (
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/configuration/validation"
	"go-with-docker-and-swagger/src/controller/model/request"
	"go-with-docker-and-swagger/src/model"
	"go-with-docker-and-swagger/src/view"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))

}
