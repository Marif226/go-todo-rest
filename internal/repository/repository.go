package repository

import (
	"github.com/Marif226/go-todo-rest/internal/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error) 
}

type TodoList interface {

}

type TodoItem interface{

}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}