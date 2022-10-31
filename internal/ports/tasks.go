package ports

import (
	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"

	"context"
)

type Tasks interface{
	CreateTask(ctx context.Context, email string, taskRequest models.TaskRequest) (string, error)
	DeleteTask(ctx context.Context, email, taskId string) error
	GetTasks(ctx context.Context, email string) (*[]models.Task, error)
	GetTaskDescription(ctx context.Context, email, taskId string) (string, error)
	ApproveTask(ctx context.Context, email, taskId string) error
	DeclineTask(ctx context.Context, email, taskId string) error
}
