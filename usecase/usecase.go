package usecase

import "github.com/dapatilham/learnapi/usecase/todolist"

type Pool struct {
	TodolistUsecase todolist.Usecase
}

func NewPoolUsecases(todolistUsecase todolist.Usecase) Pool {
	return Pool{
		TodolistUsecase: todolistUsecase,
	}
}
