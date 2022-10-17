package memory

import (
	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"
)

type MemoryStorage struct {
	storage map[string]*[]models.Task
}

func New() (*MemoryStorage, error) {
	task_storage := make(map[string]*[]models.Task, 3)
	task_storage["email1"] = &[]models.Task{}
	task_storage["email2"] = &[]models.Task{}

	return &MemoryStorage{
		storage: task_storage,
	}, nil
}
