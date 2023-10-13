package task

import (
	"GoWithMongo/datstore/entity"
	"GoWithMongo/datstore/repository"

	"go.mongodb.org/mongo-driver/bson"
)

type TaskService interface {
	CreateTask(task *entity.Task) error
	EditTask(taskID string, updatedTask *entity.Task) (*entity.Task, error)
	GetTaskByID(taskID string) (*entity.Task, error)
	MarkTaskComplete(taskID string) (*entity.Task, error)
	SearchTasks(filter bson.M) ([]*entity.Task, error) // Updated return type to []*entity.Task
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: repo,
	}
}

func (s *taskService) CreateTask(task *entity.Task) error {
	return s.taskRepo.CreateTask(*task)
}

func (s *taskService) EditTask(taskID string, updatedTask *entity.Task) (*entity.Task, error) {
	return s.taskRepo.UpdateTask(taskID, *updatedTask)
}

func (s *taskService) GetTaskByID(taskID string) (*entity.Task, error) {
	return s.taskRepo.GetTaskByID(taskID)
}

func (s *taskService) MarkTaskComplete(taskID string) (*entity.Task, error) {
	return s.taskRepo.MarkTaskComplete(taskID)
}

func (s *taskService) SearchTasks(filter bson.M) ([]*entity.Task, error) {
	tasks, err := s.taskRepo.RetrieveTasks(filter)
	if err != nil {
		return nil, err
	}

	// Convert tasks to []*entity.Task
	var tasksPtr []*entity.Task
	for _, task := range tasks {
		tasksPtr = append(tasksPtr, &task)
	}

	return tasksPtr, nil
}
