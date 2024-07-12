package controller

import (
	"go-with-docker-and-swagger/src/configuration/validation"
	"go-with-docker-and-swagger/src/controller/model/request"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

}
