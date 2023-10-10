package task_service

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type TaskController struct {
	TaskCollection *mongo.Collection
}

func (controller *TaskController) createTask(c *gin.Context) {
	var task Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := controller.TaskCollection.InsertOne(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (controller *TaskController) editTask(c *gin.Context) {
	taskID := c.Param("taskID")

	var updatedTask Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"_id": taskID}
	update := bson.M{"$set": updatedTask}

	_, err := controller.TaskCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (controller *TaskController) searchTasks(c *gin.Context) {
	// Implement task search logic based on completion status, due date, and priority
	// Admins can search for tasks
	// Respond with the matched tasks
}

func (controller *TaskController) markTaskComplete(c *gin.Context) {
	taskID := c.Param("taskID")

	filter := bson.M{"_id": taskID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err := controller.TaskCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Raise an event for the system indicating the task completion

	c.JSON(http.StatusOK, gin.H{"message": "Task marked as complete"})
}
