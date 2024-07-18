package controller

import (
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/view"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {

	logger.Info("Init findUserByID controller", zap.String("journey", "findUserByID"))

	userID := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userID); err != nil {
		logger.Error("Error trying to validate userId", err, zap.String("journey", "findUserByID"))
		errorMessage := rest_err.NewBadRequestError(
			"UserId is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userID)
	if err != nil {
		logger.Error("Error to call findUserByID services ", err, zap.String("journey", "findUserByID"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully", zap.String("journey", "findUserByID"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller", zap.String("journey", "findUserByEmail"))

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail", err, zap.String("journey", "findUserByEmail"))
		errorMessage := rest_err.NewBadRequestError(
			"UserEmail is not a valid email",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error to call findUserByEmail services ", err, zap.String("journey", "findUserByEmail"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully", zap.String("journey", "findUserByEmail"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
