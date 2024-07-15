package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {

	userID := c.Param("userId")

	fmt.Println(userID)

}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	userID := c.Param("userId")

	fmt.Println(userID)
}
