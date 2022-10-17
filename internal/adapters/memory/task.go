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

	*s.storage[email] = append(*s.storage[email], task)
	return nil
}

func (s *MemoryStorage) GetTasks(ctx context.Context, email string) (*[]models.Task, error) {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()

	tasks, ok := s.storage[email]
	if !ok {
		zapctx.Error(ctx, "user not found", zap.String("user", email))
		err := fmt.Errorf("%w: user", models.ErrNotFound)
		metrics.SetError(span, err)
		return &[]models.Task{}, err
	}
	return tasks, nil
}

func (s *MemoryStorage) GetTaskById(ctx context.Context, email, taskId string) (*models.Task, error) {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()

	tasks, err := s.GetTasks(ctx, email)
	if err != nil {
		return &models.Task{}, err
	}
	for i := range(*tasks) {
		if (*tasks)[i].Id == taskId {
			return &((*tasks)[i]), nil
		}
	}
	zapctx.Error(ctx, "task not found", zap.String("task", taskId))
	err = fmt.Errorf("%w: task", models.ErrNotFound)
	metrics.SetError(span, err)
	return &models.Task{}, err
}