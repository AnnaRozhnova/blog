package repository

import (
	"fmt"

	"github.com/AnnaRozhnova/blog"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// jjj
func (r *AuthPostgres) CreateUser(user blog.User) error {

	query := fmt.Sprintf("INSERT INTO %s values ($1, $2, $3)", usersTable)
	_, err := r.db.Exec(query, user.Username, user.Name, user.Password)
	if err != nil {
		return err
	}

	return nil
}


func (r *AuthPostgres) GetUser(username, password string) (blog.User, error) {
	var user blog.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}