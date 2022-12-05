package tests

import (
	"gitlab.com/golang-hse-2022/team1/tasks/internal/adapters/memory"
	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"
	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/usecases"

	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	
	t.Run("correct approve", func(t *testing.T) {
		memory, err := memory.New()
		assert.Nil(t, err)

		logic, err := usecases.New(memory)
		assert.Nilf(t, err, "err should be nil")
		
		authorEmail := "aidaramrin@gmail.com"

		a1 := models.Approver{
			Email: "email1",
			Status: "not set",
		}
		a2 := models.Approver{
			Email: "email2",
			Status: "not set",
		}
		a3 := models.Approver{
			Email: "email3",
			Status: "not set",
		}

		taskBody := models.TaskRequest{
			Title: "very first",
			Description: "the very first task",
			Approvers: []*models.Approver{&a1, &a2, &a3},
		}
		
		ctx := context.Background()
		taskId, err := logic.CreateTask(ctx, authorEmail, taskBody)
		assert.Nil(t, err)

		_, err = logic.ApproveTask(ctx, "email1", taskId)
		assert.Nil(t, err)
		assert.Equal(t, "approved", a1.Status)

		_, err = logic.ApproveTask(ctx, "email2", taskId)
		assert.Nil(t, err)
		assert.Equal(t, "approved", a2.Status)

		_, err = logic.ApproveTask(ctx, "email3", taskId)
		assert.Nil(t, err)
		assert.Equal(t, "approved", a3.Status)

		task, err := memory.GetTaskById(ctx, taskId)
		assert.Nil(t, err)

		assert.Equal(t, "approved", task.Status)
	})

	t.Run("decline", func(t *testing.T) {
		memory, err := memory.New()
		assert.Nil(t, err)

		logic, err := usecases.New(memory)
		assert.Nilf(t, err, "err should be nil")
		
		authorEmail := "aidaramrin@gmail.com"

		a1 := models.Approver{
			Email: "email1",
			Status: "not set",
		}
		a2 := models.Approver{
			Email: "email2",
			Status: "not set",
		}
		a3 := models.Approver{
			Email: "email3",
			Status: "not set",
		}

		taskBody := models.TaskRequest{
			Title: "very first",
			Description: "the very first task",
			Approvers: []*models.Approver{&a1, &a2, &a3},
		}
		
		ctx := context.Background()
		taskId, err := logic.CreateTask(ctx, authorEmail, taskBody)
		assert.Nil(t, err)

		_, err = logic.ApproveTask(ctx, "email1", taskId)
		assert.Nil(t, err)
		assert.Equal(t, "approved", a1.Status)

		_, err = logic.DeclineTask(ctx, "email2", taskId)
		assert.Nil(t, err)
		assert.Equal(t, "declined", a2.Status)

		task, err := memory.GetTaskById(ctx, taskId)
		assert.Nil(t, err)

		assert.Equal(t, "declined", task.Status)
	})

	t.Run("incorrect approve after decline", func(t *testing.T) {
		memory, err := memory.New()
		assert.Nil(t, err)

		logic, err := usecases.New(memory)
		assert.Nilf(t, err, "err should be nil")
		
		authorEmail := "aidaramrin@gmail.com"

		a1 := models.Approver{
			Email: "email1",
			Status: "not set",
		}
		a2 := models.Approver{
			Email: "email2",
			Status: "not set",
		}
		a3 := models.Approver{
			Email: "email3",
			Status: "not set",
		}

		taskBody := models.TaskRequest{
			Title: "very first",
			Description: "the very first task",
			Approvers: []*models.Approver{&a1, &a2, &a3},
		}
		
		ctx := context.Background()
		taskId, err := logic.CreateTask(ctx, authorEmail, taskBody)
		assert.Nil(t, err)

		_, err = logic.ApproveTask(ctx, "email1", taskId)
		assert.Nil(t, err)
		assert.Equal(t, "approved", a1.Status)
		
		_, err = logic.DeclineTask(ctx, "email2", taskId)
		assert.Nil(t, err)
		assert.Equal(t, "declined", a2.Status)

		_, err = logic.ApproveTask(ctx, "email3", taskId)
		assert.Error(t, err)

		task, err := memory.GetTaskById(ctx, taskId)
		assert.Nil(t, err)

		assert.Equal(t, "declined", task.Status)
	})
}
