package models

type TaskRequest struct {
	Title string
	Description string
	Approvers []*Approver
}
