package memory

import (
	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"
)

type MemoryStorage struct {
	tasksByAuthor map[string]*[]models.Task
	tasksById map[string]*models.Task
}

func New() (*MemoryStorage, error) {
	byAuthorStorage := make(map[string]*[]models.Task)
	byIdStorage := make(map[string]*models.Task)
	return &MemoryStorage{
		tasksByAuthor: byAuthorStorage,
		tasksById: byIdStorage,
	}, nil
}
