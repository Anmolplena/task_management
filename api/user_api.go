package api
// package user_service

// import (
// 	"context"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type UserController struct {
// 	UserCollection *mongo.Collection
// }

// func (controller *UserController) createUser(c *gin.Context) {
// 	var user User
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	_, err := controller.UserCollection.InsertOne(context.Background(), user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, user)
// }

// func (controller *UserController) deleteUser(c *gin.Context) {
// 	username := c.Param("username")

// 	_, err := controller.UserCollection.DeleteOne(context.Background(), bson.M{"username": username})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
// }

// func (controller *UserController) GetUser(c *gin.Context) {
// 	username := c.Param("username")

// 	var user User
// 	err := controller.UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }

// package user_service

// import (
// 	"GoWithMongo/middleware"

// 	"github.com/gin-gonic/gin"
// )

// func SetupUserRoutes(router *gin.Engine, controller *UserController) {
// 	router.Use(middleware.AuthMiddleware())
// 	router.POST("/users", controller.createUser)
// 	router.DELETE("/users/:username",middleware.AdminAuthMiddleware(), controller.deleteUser)
// 	router.GET("/users/:username", controller.GetUser)
// 	// Add other user-related routes
// }


