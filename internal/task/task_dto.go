package task

import (
	"GoWithMongo/datstore/entity"
)

type TaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	DueDate     string `json:"due_date"`
	Completed   bool   `json:"completed"`
	Username    string `json:"username"`
}

func NewTaskDTOFromTask(task *entity.Task) TaskDTO {
	return TaskDTO{
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Priority:    task.Priority,
		Completed:   task.Completed,
		Username: task.Username,
	}
}
