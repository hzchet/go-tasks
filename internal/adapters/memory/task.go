package memory

import (
	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"
	"gitlab.com/golang-hse-2022/team1/tasks/pkg/infra/metrics"
	"context"
	"fmt"

	"github.com/juju/zaputil/zapctx"
	"go.uber.org/zap"
)

func (s *MemoryStorage) AddTask(ctx context.Context, email string, task models.Task) error {
	_, span := metrics.FollowSpan(ctx)
	defer span.End()
	
	tasks, ok := s.tasksByAuthor[email]
	if !ok {
		s.tasksByAuthor[email] = &[]*models.Task{}
		tasks = s.tasksByAuthor[email]
	}
	*tasks = append(*tasks, &task)
	s.tasksById[task.Id] = &task
	return nil
}

func (s *MemoryStorage) GetTasks(ctx context.Context, email string) (*[]models.Task, error) {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()

	tasks, ok := s.tasksByAuthor[email]
	if !ok {
		zapctx.Error(ctx, "user not found", zap.String("user", email))
		err := fmt.Errorf("%w: user", models.ErrNotFound)
		metrics.SetError(span, err)
		return &[]models.Task{}, err
	}
	
	result := []models.Task{}
	for i := range(*tasks) {
		result = append(result, *(*tasks)[i])
	}
	return &result, nil
}

func (s *MemoryStorage) GetTaskById(ctx context.Context, taskId string) (*models.Task, error) {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()

	task, ok := s.tasksById[taskId]
	if !ok {
		zapctx.Error(ctx, "task not found", zap.String("task", taskId))
		err := fmt.Errorf("%w: task", models.ErrNotFound)
		metrics.SetError(span, err)
		return &models.Task{}, err
	}
	return task, nil
}
