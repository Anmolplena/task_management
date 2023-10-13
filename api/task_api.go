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
		taskGroup.GET("/:taskID", api.GetTaskByID)
		taskGroup.GET("/", api.SearchTasks)
		taskGroup.PATCH("/:taskID/complete", api.MarkTaskComplete)
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

func (api *TaskAPI) GetTaskByID(c *gin.Context) {
	taskID := c.Param("taskID")

	task, err := api.TaskService.GetTaskByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (api *TaskAPI) SearchTasks(c *gin.Context) {
	// Implement parsing query parameters for filtering if needed

	filter := make(map[string]interface{})
	// Example: Parse a query parameter named "status" as filter
	priority := c.Query("priority")
	completed:= c.Query("completed")
	if priority != "" {
		filter["priority"] = priority
	}
	if completed != ""{
		filter["priority"] = priority
	}

	tasks, err := api.TaskService.SearchTasks(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (api *TaskAPI) MarkTaskComplete(c *gin.Context) {
	taskID := c.Param("taskID")

	taskUpdated, err := api.TaskService.MarkTaskComplete(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, taskUpdated)
}

