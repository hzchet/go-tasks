package models

type Task struct {
	AuthorEmail string
	Body TaskRequest
	IsCancelled bool
	Id string
}
