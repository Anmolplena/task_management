package task_service

import "github.com/gin-gonic/gin"

func SetupTaskRoutes(router *gin.Engine, controller *TaskController) {
	router.POST("/tasks", controller.createTask)
	router.PUT("/tasks/:taskID", controller.editTask)
	router.GET("/tasks", controller.searchTasks)
	router.PATCH("/tasks/:taskID/complete", controller.markTaskComplete)
	// Add other task-related routes
}
