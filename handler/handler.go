package handler

import (
	"github.com/dapatilham/learnapi/handler/todolist"
	"github.com/gin-gonic/gin"
)

type Pool struct {
	TodoListHandler todolist.Handler
}

func NewPoolHandlers(engine *gin.Engine) Pool {
	return Pool{
		TodoListHandler: todolist.NewHandler(engine.Group("paper")),
	}
}
