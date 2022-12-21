package memory

import (
	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"
)

type Approver struct {
	Email    string   `bson:"email"`
	Status 	 string   `bson:"status"`
	Subtasks []string `bson:"subtasks"`
}

func (a *Approver) ToService() *models.Approver {
	return &models.Approver{
		Email:    a.Email,
		Status:   a.Status,
		Subtasks: a.Subtasks,
	}
}

func approverFromService(a *models.Approver) *Approver {
	return &Approver{
		a.Email,
		a.Status,
		a.Subtasks,
	}
}
