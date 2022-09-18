package todolist

import "github.com/dapatilham/learnapi/dto/model"

type Usecase interface {
	// create, update, delete, get
	SaveTask(req model.CreateTaskRequest) (model.CreateTaskResponse, error)
	EditTask(req model.UpdateTaskRequest) (model.UpdateTaskResponse, error)
	RemoveTask(req model.DeleteTaskRequest) (model.DeleteTaskResponse, error)
	MyTask(req model.GetTaskRequest) (model.GetTaskResponse, error)
}

type usecase struct {
}

func (u *usecase) SaveTask(req model.CreateTaskRequest) (model.CreateTaskResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *usecase) EditTask(req model.UpdateTaskRequest) (model.UpdateTaskResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *usecase) RemoveTask(req model.DeleteTaskRequest) (model.DeleteTaskResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *usecase) MyTask(req model.GetTaskRequest) (model.GetTaskResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewUsecase() Usecase {
	return &usecase{}
}
