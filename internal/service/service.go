package service

import (
	"github.com/Marif226/go-todo-rest/internal/model"
	"github.com/Marif226/go-todo-rest/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
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
		Authorization: NewAuthService(repository.Authorization),
	}
}

