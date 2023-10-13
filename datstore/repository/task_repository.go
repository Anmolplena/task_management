package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"GoWithMongo/datstore/entity"
)


type TaskRepository interface {
	CreateTask(task entity.Task) error
	UpdateTask(taskID string, updatedTask entity.Task) (*entity.Task, error)
	GetTaskByID(taskID string) (*entity.Task, error)
	MarkTaskComplete(taskID string) (*entity.Task, error)
	RetrieveTasks(filter bson.M) ([]entity.Task, error)
}

type taskRepositoryImpl struct {
	TaskCollection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) TaskRepository {
	return &taskRepositoryImpl{
		TaskCollection: collection,
	}
}

func (repo *taskRepositoryImpl) CreateTask(task entity.Task) error {
	_, err := repo.TaskCollection.InsertOne(context.Background(), task)
	return err
}

func (repo *taskRepositoryImpl) UpdateTask(taskID string, updatedTask entity.Task) (*entity.Task, error) {
	// Convert taskID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID format: %v", err)
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updatedTask}

	result := repo.TaskCollection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedTaskResult entity.Task
	if err := result.Decode(&updatedTaskResult); err != nil {
		return nil, fmt.Errorf("error updating task: %v", err)
	}

	return &updatedTaskResult, nil
}

func (repo *taskRepositoryImpl) GetTaskByID(taskID string) (*entity.Task, error) {
	// Convert taskID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID format: %v", err)
	}

	filter := bson.M{"_id": objectID}

	var task entity.Task
	err = repo.TaskCollection.FindOne(context.Background(), filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("task not found: %v", err)
		}
		return nil, fmt.Errorf("internal server error: %v", err)
	}

	return &task, nil
}

func (repo *taskRepositoryImpl) MarkTaskComplete(taskID string) (*entity.Task, error) {
	// Convert taskID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID format: %v", err)
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	result := repo.TaskCollection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedTask entity.Task
	if err := result.Decode(&updatedTask); err != nil {
		return nil, fmt.Errorf("error marking task as complete: %v", err)
	}

	return &updatedTask, nil
}

func (repo *taskRepositoryImpl) RetrieveTasks(filter bson.M) ([]entity.Task, error) {
	// Define options to sort the results by due_date and priority
	sortOptions := options.Find().SetSort(bson.D{{"due_date", 1}, {"priority", 1}})

	// Find documents in the collection that match the filter
	cursor, err := repo.TaskCollection.Find(context.Background(), filter, sortOptions)
	if err != nil {
		return nil, fmt.Errorf("error finding tasks: %v", err)
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and decode each document into a Task
	var tasks []entity.Task
	for cursor.Next(context.Background()) {
		var task entity.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, fmt.Errorf("error decoding task: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
