package user_service

import "github.com/gin-gonic/gin"

func SetupUserRoutes(router *gin.Engine, controller *UserController) {
	router.POST("/users", controller.createUser)
	router.DELETE("/users/:username", controller.deleteUser)
	// Add other user-related routes
}
