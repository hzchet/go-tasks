package usecases

import (
	"context"
	"fmt"

	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"
	"gitlab.com/golang-hse-2022/team1/tasks/internal/ports"
	"gitlab.com/golang-hse-2022/team1/tasks/pkg/infra/metrics"

	"github.com/satori/go.uuid"
)

type Tasks struct {
	taskStorage ports.TaskStorage
}

func New(taskStorage ports.TaskStorage) (*Tasks, error) {
	return &Tasks{
		taskStorage: taskStorage,
	}, nil
}

func (t *Tasks) CreateTask(ctx context.Context, email string, taskRequest models.TaskRequest) error {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()

	task := models.Task{
		AuthorEmail: email,
		Body: taskRequest,
		IsCancelled: false,
		Id: uuid.NewV1().String(),
	}
	return t.taskStorage.AddTask(ctx, email, task)
}

func (t *Tasks) DeleteTask(ctx context.Context, email, taskId string) error {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()
	
	task, err := t.taskStorage.GetTaskById(ctx, email, taskId)
	if err != nil {
		return err
	}
	task.IsCancelled = true
	return nil
}

func (t *Tasks) GetTasks(ctx context.Context, email string) ([]models.Task, error) {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()
	
	return t.GetTasks(ctx, email)
}

func (t *Tasks) GetTaskDescription(ctx context.Context, email, taskId string) (string, error) {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()
	
	task, err := t.taskStorage.GetTaskById(ctx, email, taskId)
	if err != nil {
		return "", err
	}
	return task.Body.Description, nil
}

func (t *Tasks) ApproveTask(ctx context.Context,email, taskId string) error {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()
	
	task, err := t.taskStorage.GetTaskById(ctx, email, taskId)
	if err != nil {
		return err
	}
	for i := range(task.Body.Approvers) {
		if task.Body.Approvers[i].Email == email {
			task.Body.Approvers[i].Status = "approved"
			return nil
		}
	}
	return fmt.Errorf("approver not found: %w", models.ErrNotFound)
}

func (t *Tasks) DeclineTask(ctx context.Context, email, taskId string) error {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()
	
	task, err := t.taskStorage.GetTaskById(ctx, email, taskId)
	if err != nil {
		return err
	}
	for i := range(task.Body.Approvers) {
		if task.Body.Approvers[i].Email == email {
			task.Body.Approvers[i].Status = "declined"
			return nil
		}
	}
	return fmt.Errorf("approver not found: %w", models.ErrNotFound)
}
