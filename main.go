package main

import (
	"github.com/dapatilham/learnapi/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	// ToDoList
	// postman request

	engine := gin.New()
	handler.NewPoolHandlers(engine)

	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
