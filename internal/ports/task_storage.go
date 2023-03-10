package ports

import (
	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"
	"context"
)

type TaskStorage interface {
	AddTask(ctx context.Context, task models.Task) error
	GetTasks(ctx context.Context, email string) (*[]models.Task, error)
	GetTaskById(ctx context.Context, taskId string) (*models.Task, error)
}
