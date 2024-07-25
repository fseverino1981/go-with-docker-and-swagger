package controller

import (
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// DeleteUser deletes a user with the specified ID.
// @Summary Delete User
// @Description Deletes a user based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be deleted"
// @Success 200
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /deleteUser/{userId} [delete]
func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller", zap.String("journey", "deleteUser"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
		return
	}

	err := uc.service.DeleteUserServices(userId)
	if err != nil {
		logger.Error("Error trying to call deleteUser service", err, zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"DeleteUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	c.Status(http.StatusOK)

}
