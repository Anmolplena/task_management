package main

import (
	"context"
	"log"
	"GoWithMongo/task_service"
	"GoWithMongo/user_service"

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

	userCollection := client.Database("task_management").Collection("users")
	taskCollection := client.Database("task_management").Collection("tasks")

	userController := &user_service.UserController{UserCollection: userCollection}
	taskController := &task_service.TaskController{TaskCollection: taskCollection}

	user_service.SetupUserRoutes(r, userController)
	task_service.SetupTaskRoutes(r, taskController)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
