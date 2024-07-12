package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context){
	
	userID := c.Param("userId")

	fmt.Println(userID)

}


func FindUserByEmail(c *gin.Context){
	userID := c.Param("userId")

	fmt.Println(userID)
}