package memory

import (
	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"
	"gitlab.com/golang-hse-2022/team1/tasks/pkg/infra/metrics"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Task struct {
	AuthorEmail string      `bson:"author_email`
	Body        TaskRequest `bson:"body"`
	IsCancelled bool        `bson:"is_cancelled`
	Status      string      `bson:"status"`   
	Id          string      `bson:"_id"`
}

func (t *Task) ToService() models.Task {
	return models.Task{
		AuthorEmail: t.AuthorEmail,
		Body: 		 t.Body.ToService(),
		IsCancelled: t.IsCancelled,
		Status:      t.Status,
		Id: 		 t.Id,
	}
}

func taskFromService(t *models.Task) Task {
	return Task{
		AuthorEmail: t.AuthorEmail,
		Body: 		 taskRequestFromService(&t.Body),
		IsCancelled: t.IsCancelled,
		Status: 	 t.Status,
		Id: 		 t.Id,
	}
}

func (d *Database) AddTask(ctx context.Context, task models.Task) error {
	_, span := metrics.FollowSpan(ctx)
	defer span.End()
	
	_, err := d.taskCollection.InsertOne(ctx, taskFromService(&task))
	if err != nil {
		return fmt.Errorf("cannot insert task: %w", err)
	}

	return nil
}

func (d *Database) GetTasks(ctx context.Context, email string) (*[]models.Task, error) {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()

	var tasks []models.Task

	cursor, err := d.taskCollection.Find(ctx, bson.M{"email": email})
	if err != nil {
		return &[]models.Task{}, fmt.Errorf("cannot get tasks: %w", err)
	}

	for cursor.Next(ctx) {
		var t Task
		if err := cursor.Decode(&t); err != nil {
			return &[]models.Task{}, fmt.Errorf("cannot get tasks: %w, err")
		}
		tasks = append(tasks, t.ToService())
	}

	return &tasks, nil
}

func (d *Database) GetTaskById(ctx context.Context, taskId string) (*models.Task, error) {
	ctx, span := metrics.FollowSpan(ctx)
	defer span.End()

	var task Task

	err := d.taskCollection.FindOne(ctx, bson.M{"_id": taskId}).Decode(&task)
	if err != nil {
		return &models.Task{}, fmt.Errorf("cannot get task by id: %w", err)
	}

	var t models.Task = task.ToService()
	return &t, nil
}