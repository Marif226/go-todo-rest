package repository

import (
	"fmt"

	"github.com/Marif226/go-todo-rest/internal/model"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateUser(user model.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id;", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	
	return id, nil
}