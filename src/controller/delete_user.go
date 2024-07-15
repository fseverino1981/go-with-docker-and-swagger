package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	userID := c.Param("userId")

	fmt.Println(userID)

}
