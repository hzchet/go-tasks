package memory

import (
	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"
)

type TaskRequest struct {
	Title       string      `bson:"title"`
	Description string      `bson:"descriptoin"`
	Approvers   []*Approver `bson:"approvers"`
}

func (tr *TaskRequest) ToService() models.TaskRequest {
	approvers := make([]*models.Approver, len(tr.Approvers))
	for i := range(tr.Approvers) {
		approvers[i] = tr.Approvers[i].ToService()
	}

	return models.TaskRequest{
		Title: tr.Title,
		Description: tr.Description,
		Approvers: approvers,
	}
}

func taskRequestFromService(tr *models.TaskRequest) TaskRequest {
	approvers := make([]*Approver, len(tr.Approvers))
	for i := range(tr.Approvers) {
		approvers[i] = approverFromService(tr.Approvers[i])
	}
	
	return TaskRequest{
		Title:     	 tr.Title,
		Description: tr.Description,
		Approvers:	 approvers,
	}
}
