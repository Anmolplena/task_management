package task_service

import (
	"GoWithMongo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine, controller *TaskController) {
	router.Use(middleware.AuthMiddleware())
	router.POST("/tasks", middleware.AdminAuthMiddleware(), controller.createTask)
	router.PUT("/tasks/:taskID", middleware.AdminAuthMiddleware(), controller.editTask)
	router.GET("/tasks/:taskID", controller.GetTaskById)
	router.GET("/tasks", middleware.AdminAuthMiddleware(), controller.searchTasks)
	router.PATCH("/tasks/:taskID/complete", controller.markTaskComplete)
	// Add other task-related routes
}
