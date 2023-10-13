package main

import (
	"context"
	"log"

	"GoWithMongo/api"
	"GoWithMongo/datstore/repository"
	"GoWithMongo/internal/task"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	taskCollection := client.Database("task_management").Collection("tasks")

	// Create the task repository
	taskRepo := repository.NewTaskRepository(taskCollection)

	// Set up task service using the task repository
	taskService := task.NewTaskService(taskRepo)

	// Set up task API
	api.NewTaskAPI(r, taskService)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
