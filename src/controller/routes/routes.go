package routes

import (
	"go-with-docker-and-swagger/src/controller"
	"go-with-docker-and-swagger/src/model"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.VerifTokenMiddleware, userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", model.VerifTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", model.VerifTokenMiddleware, userController.CreateUser)
	r.PUT("/updateUser/:userId", model.VerifTokenMiddleware, userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", model.VerifTokenMiddleware, userController.DeleteUser)

	r.POST("/login", userController.LoginUser)

}
