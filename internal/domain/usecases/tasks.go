package usecases

import (
	"context"
	"fmt"
	"log"

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

func (t *Tasks) CreateTask(ctx context.Context, email string, taskRequest models.TaskRequest) (string, error) {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()

	taskId := uuid.NewV1().String()
	task := models.Task{
		AuthorEmail: email,
		Body: taskRequest,
		IsCancelled: false,
		Id: taskId,
	}

	log.Printf("send mail to %s: approve/decline new task %s\n", task.Body.Approvers[0].Email, taskId)
	return taskId, t.taskStorage.AddTask(ctx, email, task)
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

func (t *Tasks) GetTasks(ctx context.Context, email string) (*[]models.Task, error) {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()
	
	return t.taskStorage.GetTasks(ctx, email)
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

func sendResult(ctx context.Context, result string, taskId string, approvers []*models.Approver) error {
	for i := range(approvers) {
		log.Printf("send mail to %s: task %s, status: %s\n", approvers[i].Email, taskId, result)
	}
	return nil
}

func (t *Tasks) ApproveTask(ctx context.Context, email, taskId string) error {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()
	
	task, err := t.taskStorage.GetTaskById(ctx, email, taskId)
	if err != nil {
		return err
	}
	var approverIndex int = -1
	for i := range(task.Body.Approvers) {
		if task.Body.Approvers[i].Email == email {
			task.Body.Approvers[i].Status = "approved"
			approverIndex = i
			break
		}
	}
	if approverIndex != -1 && approverIndex + 1 < len(task.Body.Approvers) {
		log.Printf("send mail to %s: approve/decline new task %s\n", task.Body.Approvers[approverIndex + 1].Email, taskId)
	} else if approverIndex == len(task.Body.Approvers) - 1 {
		return sendResult(ctx, "approved", taskId, task.Body.Approvers)
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
			return sendResult(ctx, "declined", taskId, task.Body.Approvers)
		}
	}
	return fmt.Errorf("approver not found: %w", models.ErrNotFound)
}
