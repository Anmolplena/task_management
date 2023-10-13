// task_api.go
package api

import (
	"GoWithMongo/datstore/entity"
	"GoWithMongo/internal/task"
	"GoWithMongo/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskAPI struct {
	TaskService task.TaskService
}

func NewTaskAPI(router *gin.Engine, taskService task.TaskService) {
	api := TaskAPI{
		TaskService: taskService,
	}

	taskGroup := router.Group("/tasks")
	taskGroup.Use(middleware.AuthMiddleware())
	{
		taskGroup.POST("/", api.CreateTask)
		taskGroup.PUT("/:taskID", api.EditTask)
		// taskGroup.GET("/:taskID", api.GetTaskByID)
		// taskGroup.GET("/", api.SearchTasks)
		// taskGroup.PATCH("/:taskID/complete", api.MarkTaskComplete)
	}
}

func (api *TaskAPI) CreateTask(c *gin.Context) {
	var task entity.Task

	// Bind the request body to the task struct
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the CreateTask function with the deserialized task
	err := api.TaskService.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (api *TaskAPI) EditTask(c *gin.Context) {
	taskID := c.Param("taskID")

	var updatedTask entity.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, taskUpdated := api.TaskService.EditTask(taskID, &updatedTask)

	c.JSON(http.StatusOK, taskUpdated)
}

// func (api *TaskAPI) GetTaskByID(c *gin.Context) {
// 	taskID := c.Param("taskID")

// 	// Call the GetTaskByID function in the service
// 	// Implement the logic to call the service's GetTaskByID method
// 	// task, err := api.TaskService.GetTaskByID(taskID)
// 	// ...

// 	// Placeholder response for now
// 	c.JSON(http.StatusOK, gin.H{"message": "Task retrieved successfully"})
// }

// func (api *TaskAPI) SearchTasks(c *gin.Context) {
// 	// Parse query parameters for filtering, if needed
// 	// ...

// 	// Call the SearchTasks function in the service
// 	// Implement the logic to call the service's SearchTasks method
// 	// tasks, err := api.TaskService.SearchTasks(filter)
// 	// ...

// 	// Placeholder response for now
// 	c.JSON(http.StatusOK, gin.H{"message": "Tasks retrieved successfully"})
// }

// func (api *TaskAPI) MarkTaskComplete(c *gin.Context) {
// 	taskID := c.Param("taskID")

// 	// Call the MarkTaskComplete function in the service
// 	// Implement the logic to call the service's MarkTaskComplete method
// 	// updatedTask, err := api.TaskService.MarkTaskComplete(taskID)
// 	// ...

// 	// Placeholder response for now
// 	c.JSON(http.StatusOK, gin.H{"message": "Task marked as complete"})
// }

// Add other task-related functions as needed
