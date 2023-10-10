package task_service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// Convert taskID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updatedTask}

	result := controller.TaskCollection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedTaskResult Task
	if err := result.Decode(&updatedTaskResult); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "updatedTask": updatedTaskResult})
}

func (controller *TaskController) searchTasks(c *gin.Context) {

	// Parse query parameters for sorting
	completedStatus := c.Query("completed")
	dueDate := c.Query("due_date")
	priority := c.Query("priority")

	// Prepare the filter based on query parameters
	filter := bson.M{}
	if completedStatus != "" {
		completed, err := strconv.ParseBool(completedStatus)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for completed"})
			return
		}
		filter["completed"] = completed
	}

	if dueDate != "" {
		filter["due_date"] = dueDate
	}

	if priority != "" {
		filter["priority"] = priority
	}

	// Retrieve tasks based on the filter
	tasks, err := controller.retrieveTasks(filter)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (controller *TaskController) retrieveTasks(filter bson.M) ([]Task, error) {
	// Define options to sort the results by due_date and priority
	sortOptions := options.Find().SetSort([]bson.E{{Key: "due_date", Value: 1}, {Key: "priority", Value: 1}})

	// Find documents in the collection that match the filter
	cursor, err := controller.TaskCollection.Find(context.Background(), filter, sortOptions)
	if err != nil {
		fmt.Println("Error finding tasks:", err)  // Print the error
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and decode each document into a Task
	var tasks []Task
	for cursor.Next(context.Background()) {
		var task Task
		if err := cursor.Decode(&task); err != nil {
			fmt.Println("Error decoding task:", err)  // Print the error
			return nil, err
		}
		tasks = append(tasks, task)
	}

	fmt.Println("Retrieved tasks:", tasks)  // Print the retrieved tasks
	return tasks, nil
}





func (controller *TaskController) GetTaskById(c *gin.Context) {
	taskID := c.Param("taskID")

	// Convert taskID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}

	filter := bson.M{"_id": objectID}

	var task Task
	err = controller.TaskCollection.FindOne(context.Background(), filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (controller *TaskController) markTaskComplete(c *gin.Context) {
	taskID := c.Param("taskID")

	// Convert taskID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	result := controller.TaskCollection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedTask Task
	if err := result.Decode(&updatedTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marking task as complete"})
		return
	}

	// Raise an event for the system indicating the task completion

	c.JSON(http.StatusOK, gin.H{"message": "Task marked as complete", "updatedTask": updatedTask})
}
