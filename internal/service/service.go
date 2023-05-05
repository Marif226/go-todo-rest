package service

import "github.com/Marif226/go-todo-rest/internal/repository"

type Authorization interface {

}

type TodoList interface {

}

type TodoItem interface{

}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func New(repository *repository.Repository) *Service {
	return &Service{

	}
}