package repository

import "github.com/dapatilham/learnapi/repository/todolist"

type Pool struct {
	TodolistRepository todolist.Repository
}

func NewPoolRepositories(todolistRepository todolist.Repository) Pool {
	return Pool{
		TodolistRepository: todolistRepository,
	}
}
