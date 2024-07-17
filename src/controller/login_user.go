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

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller", zap.String("journey", "loginUser"))

	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying validate user info", err, zap.String("journey", "loginUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)

	domainResult, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call loginUser service", err, zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"LoginUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))

}
