package todolist

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateTask()
	UpdateTask()
	DeleteTask()
	GetTasks()
}

type handler struct {
	group *gin.RouterGroup
}

func (h *handler) CreateTask() {
	const endpoint = `/task/create`
	h.group.POST(endpoint, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "create endpoint ok"})
	})
}

func (h *handler) UpdateTask() {
	const endpoint = `/task/update`
	h.group.PUT(endpoint, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "update endpoint ok"})
	})
}

func (h *handler) DeleteTask() {
	const endpoint = `/task/delete`
	h.group.DELETE(endpoint, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "delete endpoint ok"})
	})
}

func (h *handler) GetTasks() {
	const endpoint = `/task/list`
	h.group.GET(endpoint, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "list endpoint ok"})
	})
}

func NewHandler(group *gin.RouterGroup) Handler {
	h := &handler{
		group: group,
	}
	h.CreateTask()
	h.DeleteTask()
	h.UpdateTask()
	h.GetTasks()
	return h
}
