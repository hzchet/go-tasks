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
		Status: "in consideration",
		Id: taskId,
	}
	
	err := t.taskStorage.AddTask(ctx, task)

	if len(task.Body.Approvers) > 0 {
		log.Printf("send mail to %s: approve/decline new task %s\n", task.Body.Approvers[0].Email, taskId)
	}

	return taskId, err
}

func (t *Tasks) DeleteTask(ctx context.Context, email, taskId string) error {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()
	
	task, err := t.taskStorage.GetTaskById(ctx, taskId)
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
	
	task, err := t.taskStorage.GetTaskById(ctx, taskId)
	if err != nil {
		return "", err
	}

	return task.Body.Description, nil
}

func (t *Tasks) ApproveTask(ctx context.Context, email, taskId string) (models.Message, error) {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()
	
	task, err := t.taskStorage.GetTaskById(ctx, taskId)
	if err != nil {
		return models.Message{}, err
	}
	if task.Status == "declined" {
		return models.Message{}, fmt.Errorf("task already declined: %w", models.ErrForbidden)
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
		message := models.Message{
			Theme: "You assigned as an approver!",
			Recievers: []*models.Approver{task.Body.Approvers[approverIndex + 1]},
			Body: "Approve/decline new task with id",
		}
		return message, nil
	} else if approverIndex == len(task.Body.Approvers) - 1 {
		task.Status = "approved"
		message := models.Message{
			Theme: "Task is approved!",
			Recievers: task.Body.Approvers,
			Body: "Task is approved by all approvers!!",
		}
		return message, nil
	}
	return models.Message{}, fmt.Errorf("approver not found: %w", models.ErrNotFound)
}

func (t *Tasks) DeclineTask(ctx context.Context, email, taskId string) (models.Message, error) {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()
	
	task, err := t.taskStorage.GetTaskById(ctx, taskId)
	if err != nil {
		return models.Message{}, err
	}
	if task.Status == "approved" {
		return models.Message{}, fmt.Errorf("task already approved: %w", models.ErrForbidden)
	}
	for i := range(task.Body.Approvers) {
		if task.Body.Approvers[i].Email == email {
			task.Body.Approvers[i].Status = "declined"
			task.Status = "declined"
			message := models.Message{
				Theme: "Task is declined!",
				Recievers: task.Body.Approvers,
				Body: "Task is declined by one of the approvers!!",
			}
			return message, nil
		}
	}
	return models.Message{}, fmt.Errorf("approver not found: %w", models.ErrNotFound)
}
