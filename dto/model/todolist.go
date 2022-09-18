package model

import (
	"github.com/dapatilham/learnapi/dto/entity"
	"github.com/dapatilham/learnapi/dto/standard"
)

// creation task request response object
type (
	CreateTaskRequest struct {
		UserID   string `json:"userId"`
		Title    string `json:"tittle"`
		Desc     string `json:"desc"`
		Done     bool   `json:"done"`
		Category string `json:"category"`
		Priority int    `json:"priority"`
		DueDate  string `json`
	}
	CreateTaskResponse struct {
		standard.Response
	}
)

// Modification  task request response object
type (
	UpdateTaskRequest struct {
		ID       string `json:"id"`
		Title    string `json:"tittle"`
		Desc     string `json:"desc"`
		Done     bool   `json:"done"`
		Category string `json:"category"`
		Priority int    `json:"priority"`
		DueDate  string `json`
	}
	UpdateTaskResponse struct {
		standard.Response
	}
)

// deletion task request response object
type (
	DeleteTaskRequest struct {
		ID string `json:"id"`
	}
	DeleteTaskResponse struct {
		standard.Response
	}
)

// retrieval task request response object
type (
	GetTaskRequest struct {
		UserID string `json:"userId"`
	}
	GetTaskResponse struct {
		standard.Response
		List []entity.TodoList `json:"list"`
	}
)
