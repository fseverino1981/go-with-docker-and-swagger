package routes

import (
	"go-with-docker-and-swagger/src/controller"
	"go-with-docker-and-swagger/src/model"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.VerifTokenMiddleware, userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", model.VerifTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", model.VerifTokenMiddleware, userController.CreateUser)
	r.PUT("/updateUser/:userId", model.VerifTokenMiddleware, userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", model.VerifTokenMiddleware, userController.DeleteUser)

	r.POST("/login", userController.LoginUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
