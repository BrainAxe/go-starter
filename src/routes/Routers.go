package routes

import (
	"go-starter/src/controller"
	"go-starter/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/api/v1")
	{
		userGroup.POST("/login", controller.LoginController)
		userGroup.POST("/token/refresh", controller.RefreshTokenController)

		userGroup.POST("/registration", controller.CreateUser)
		userGroup.GET("/users", middleware.TokenAuthMiddleware(), controller.GetUserList)
		userGroup.PUT("/users/:id", controller.UpdateUser)
		userGroup.DELETE("/users/:id", controller.DeleteUserById)

	}
	return r
}
