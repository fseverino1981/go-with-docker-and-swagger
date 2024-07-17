package controller

import (
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/configuration/validation"
	"go-with-docker-and-swagger/src/controller/model/request"
	"go-with-docker-and-swagger/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller", zap.String("journey", "updateUser"))
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying validate user info", err, zap.String("journey", "updateUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
		return
	}
	domain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Age)

	err := uc.service.UpdateUserServices(userId, domain)
	if err != nil {
		logger.Error("Error trying to call updateUser service", err, zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"UpdateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)

}
