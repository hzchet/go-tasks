package http

import (
	"encoding/json"
	"net/http"

	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

// Approve
// @ID Approve
// @tags handler
// @Summary Approve task by taskId
// @Description Approve task by taskId in case it has been assigned to the user
// @Success 200
// @Failure 400 {string} string "bad request"
// @Failure 403 {string} string "forbidden"
// @Failure 500 {string} string "internal error"
// @Router /approve/:taskId [post]
func (a *Adapter) approve(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "email not set in context",
		})
	}

	message, err := a.tasks.ApproveTask(ctx, email.(string), taskId)
	if err != nil {
		a.BindError(ctx, err)		
	}
	err = a.writer.WriteMessages(ctx, kafka.Message{})
	if err != nil {
		a.BindError(ctx, err)
	}
	json_message, err := json.Marshal(&message)
	if err != nil {
		a.BindError(ctx, err)
	}
	a.writer.WriteMessages(ctx, kafka.Message{Value: []byte(json_message)})
}

// Decline
// @ID Decline
// @tags handler
// @Summary Decline task by taskId
// @Description Decline task by taskId in case it has been assigned to the user
// @Success 200
// @Failure 400 {string} string "bad request"
// @Failure 403 {string} string "forbidden"
// @Failure 500 {string} string "internal error"
// @Router /decline/:taskId [post]
func (a *Adapter) decline(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "email not set in context",
		})
	}

	message, err := a.tasks.DeclineTask(ctx, email.(string), taskId)
	if err != nil {
		a.BindError(ctx, err)		
	}
	err = a.writer.WriteMessages(ctx, kafka.Message{})
	if err != nil {
		a.BindError(ctx, err)
	}
	json_message, err := json.Marshal(&message)
	if err != nil {
		a.BindError(ctx, err)
	}
	a.writer.WriteMessages(ctx, kafka.Message{Value: []byte(json_message)})
}

// Delete
// @ID Delete
// @tags handler
// @Summary Delete task by taskId
// @Description Delete task by taskId in case it has been assigned to the user
// @Success 200
// @Failure 400 {string} string "bad request"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal error"
// @Router /:taskId [delete]
func (a *Adapter) delete(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "email not set in context",
		})
	}

	if err := a.tasks.DeleteTask(ctx, email.(string), taskId); err != nil {
		a.BindError(ctx, err)
	}
}

// GetTasks
// @ID GetTasks
// @tags handler
// @Summary get all the tasks
// @Description get tasks created by the user
// @Success 200 
// @Failure 400 {string} string "bad request"
// @Failure 403 {string} string "forbidden"
// @Failure 500 {string} string "internal error"
// @Router / [get]
func (a *Adapter) getTasks(ctx *gin.Context) {
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "email not set in context",
		})
	}

	tasks, err := a.tasks.GetTasks(ctx, email.(string))
	if err != nil {
		a.BindError(ctx, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"tasks": *tasks,
	})
}

// GetDescription
// @ID GetDescription
// @tags handler
// @Summary get task description by the taskId
// @Description get task description by the taskId
// @Success 200 {string} string
// @Failure 400 {string} string "bad request"
// @Failure 403 {string} string "forbidden"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal error"
// @Router /:taskId [get]
func (a *Adapter) getDescription(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "email not set in context",
		})
	}

	desc, err := a.tasks.GetTaskDescription(ctx, email.(string), taskId)
	if err != nil {
		a.BindError(ctx, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"taskDescription": desc,
	})
}

// Create
// @ID Create
// @tags handler
// @Summary create task
// @Description create task by the given taskRequest and return taskId
// @Success 201 {string} string
// @Failure 400 {string} string "bad request"
// @Failure 403 {string} string "forbidden"
// @Failure 500 {string} string "internal error"
// @Router / [post]
func (a *Adapter) create(ctx *gin.Context) {	
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "email not set in context",
		})
	}
	
	taskRequest := models.TaskRequest{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(&taskRequest); err != nil {
		a.BindError(ctx, err)
	}
	
	taskId, err := a.tasks.CreateTask(ctx, email.(string), taskRequest)
	if err != nil {
		a.BindError(ctx, err)
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"taskId": taskId,
	})
}
