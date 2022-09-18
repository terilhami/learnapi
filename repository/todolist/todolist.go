package todolist

import (
	"time"

	"github.com/dapatilham/learnapi/dto/model/entity"
)

type Repository interface {
	WriteTask(data entity.TodoList) error               // create
	UpdateTask(data entity.TodoList) error              // update
	EraseTask(data entity.TodoList) error               // delete
	ReadTasks(userId string) ([]entity.TodoList, error) // Get

}

type repository struct {
}

func NewPoolRepository() Repository {
	return &repository{}
}

func (r *repository) WriteTask(data entity.TodoList) error {
	return nil
}

func (r *repository) UpdateTask(data entity.TodoList) error {
	return nil
}

func (r *repository) EraseTask(data entity.TodoList) error {
	return nil
}

func (r *repository) ReadTasks(userId string) ([]entity.TodoList, error) {
	return []entity.TodoList{
		{
			ID:        "111",
			Title:     "TES",
			Desc:      "TES",
			Done:      true,
			Category:  "public",
			Priority:  1,
			DueDate:   "asap",
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}, nil
}

func NewRepository() Repository {
	return &repository{}
}
