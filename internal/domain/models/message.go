package models

type Message struct {
	Theme string
	Recievers []*Approver
	Body string
}
