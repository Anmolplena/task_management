package user_service

import (
	"GoWithMongo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, controller *UserController) {
	router.Use(middleware.AuthMiddleware())
	router.POST("/users", controller.createUser)
	router.DELETE("/users/:username",middleware.AdminAuthMiddleware(), controller.deleteUser)
	router.GET("/users/:username", controller.GetUser)
	// Add other user-related routes
}
