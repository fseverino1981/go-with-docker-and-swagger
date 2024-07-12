package controller

import (
	"fmt"
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/controller/model/request"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	userID := c.Param("userId")

	fmt.Println(userID)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, erro=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
}
