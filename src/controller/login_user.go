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

// LoginUser allows a user to log in and obtain an authentication token.
// @Summary User Login
// @Description Allows a user to log in and receive an authentication token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userLogin body request.UserLogin true "User login credentials"
// @Success 200 {object} response.UserResponse "Login successful, authentication token provided"
// @Header 200 {string} Authorization "Authentication token"
// @Failure 403 {object} rest_err.RestErr "Error: Invalid login credentials"
// @Router /login [post]
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

	domainResult, token, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call loginUser service", err, zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"LoginUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"))

	c.Header("Authorization", token)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))

}
