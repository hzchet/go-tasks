package tests

import (
	"gitlab.com/golang-hse-2022/team1/tasks/internal/application"
	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"
	"gitlab.com/golang-hse-2022/team1/tasks/pkg/infra/logger"

	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
	"encoding/base64"

	"github.com/stretchr/testify/suite"
)

type integraTestSuite struct {
	suite.Suite
	app *application.App
	tasksClient *http.Client
	accessCookie *http.Cookie
	refreshCookie *http.Cookie
	withAuth bool  // use auth or mock instead
}

func TestIntegraTestSuite(t *testing.T) {
	suite.Run(t, &integraTestSuite{})
}

func (s *integraTestSuite) login() {
	client := &http.Client{Timeout: 10 * time.Second}
	loginEndpoint := "http://localhost:8080/auth/api/v1/login"
	req, err := http.NewRequest(http.MethodPost, loginEndpoint, nil)
	
	if err != nil {
		log.Fatal("login failed")
	}
	
	email := "email1"
	password := "password1"
	str := fmt.Sprintf("%s:%s", email, password)
	str = base64.StdEncoding.EncodeToString([]byte(str))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", str))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("login failed")
	}
	defer resp.Body.Close()

	cookies := resp.Cookies()
	if len(cookies) != 2 {
		fmt.Printf("\n\nlen(cookies) = %d\n\n", len(cookies))
		log.Fatal("\n\nBRUH!!!!\n\n")
	}
	s.accessCookie = cookies[0]
	s.refreshCookie = cookies[1]
}

func (s *integraTestSuite) SetupSuite() {
	l, err := logger.New()
	if err != nil {
		log.Fatalf("logger initialization failed: %s", err.Error())
	}
	app := application.New(l)
	s.app = app
	s.withAuth = true
	err = s.app.Start(s.withAuth)
	if err != nil {
		l.Sugar().Fatalf("app not started: %s", err.Error())
	}
	s.tasksClient = &http.Client{Timeout: 10 * time.Second}
	if s.withAuth {
		s.login()
	}
}

func (s *integraTestSuite) TearDownSuite() {
	s.app.Stop(context.Background())
}

type TaskIdProjection struct {
	TaskId string
}

type DescriptionProjection struct {
	TaskDescription string
}

type TasksProjection struct {
	Tasks []models.Task
}

func (s *integraTestSuite) TestCreateTask() {
	endpoint := "http://localhost:3000/tasks/api/v1/tasks/"

	var jsonData = []byte(`{
		"title": "test_create",
		"description": "meow meow",
		"approvers": [
			{
				"email": "aamrin@edu.hse.ru",
				"status": "not set"
			},
			{
				"email": "aidar.amrin@gmail.com",
				"status": "not set"
			},
			{
				"email": "aidaramrin17052003@gmail.com",
				"status": "not set"
			}
		]
	}`)

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	s.NoError(err)
	if s.withAuth {
		req.AddCookie(s.accessCookie)
		req.AddCookie(s.refreshCookie)
	}
	respCreate, err := s.tasksClient.Do(req)
	s.NoError(err)
	defer respCreate.Body.Close()

	s.Equal(http.StatusCreated, respCreate.StatusCode)
	
	var id TaskIdProjection
	s.NoError(json.NewDecoder(respCreate.Body).Decode(&id))
	
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", endpoint, id.TaskId), nil)
	s.NoError(err)
	if s.withAuth {
		req.AddCookie(s.accessCookie)
		req.AddCookie(s.refreshCookie)
	}
	respGet, err := s.tasksClient.Do(req)
	s.NoError(err)
	defer respGet.Body.Close()
	
	var description DescriptionProjection
	s.NoError(json.NewDecoder(respGet.Body).Decode(&description))

	s.Equal("meow meow", description.TaskDescription)
}

func (s *integraTestSuite) TestDeleteTask() {
	endpoint := "http://localhost:3000/tasks/api/v1/tasks/"

	var jsonData = []byte(`{
		"title": "test_delete",
		"description": "bark bark",
		"approvers": [
			{
				"email": "superoriginalemail@gmail.com",
				"status": "not set"
			},
			{
				"email": "aidar.amrin@gmail.com",
				"status": "not set"
			},
			{
				"email": "aidaramrin17052003@gmail.com",
				"status": "not set"
			}
		]
	}`)

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	s.NoError(err)
	if s.withAuth {
		req.AddCookie(s.accessCookie)
		req.AddCookie(s.refreshCookie)
	}
	respCreate, err := s.tasksClient.Do(req)
	s.NoError(err)
	defer respCreate.Body.Close()
	s.Equal(http.StatusCreated, respCreate.StatusCode)
	
	var id TaskIdProjection
	s.NoError(json.NewDecoder(respCreate.Body).Decode(&id))

	req, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", endpoint, id.TaskId), nil)
	s.NoError(err)
	if s.withAuth {
		req.AddCookie(s.accessCookie)
		req.AddCookie(s.refreshCookie)
	}
	respDelete, err := s.tasksClient.Do(req)
	s.NoError(err)
	defer respDelete.Body.Close()
	s.Equal(http.StatusOK, respDelete.StatusCode)
	
	req, err = http.NewRequest(http.MethodGet, endpoint, nil)
	s.NoError(err)
	if s.withAuth {
		req.AddCookie(s.accessCookie)
		req.AddCookie(s.refreshCookie)
	}
	respGet, err := s.tasksClient.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, respDelete.StatusCode)
	defer respGet.Body.Close()

	var proj TasksProjection
	s.NoError(json.NewDecoder(respGet.Body).Decode(&proj))
	s.Greater(len(proj.Tasks), 0)
	s.Equal(true, proj.Tasks[len(proj.Tasks) - 1].IsCancelled)
}
