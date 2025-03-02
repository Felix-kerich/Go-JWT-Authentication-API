package routes

import (
	"auth_api/handlers"
	"auth_api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.RegisterUser)
		auth.GET("/verify-email", authHandler.VerifyEmail)
		auth.POST("/login", authHandler.LoginUser)
	}

	user := r.Group("/users")
	user.Use(middlewares.AuthMiddleware())
	{
		user.GET("/", userHandler.GetUsers)
		user.GET("/:id", userHandler.GetUserByID)
		user.PUT("/:id", userHandler.UpdateUser)
		user.DELETE("/:id", userHandler.DeleteUser)
		user.POST("/:id/recover", userHandler.RecoverUser)
	}

	return r
}
