package repository

import (
	"fmt"

	"github.com/AnnaRozhnova/blog"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GeyAll() ([]blog.User, error) {
	var users []blog.User
	query := fmt.Sprintf("SELECT username, name FROM %s", usersTable)
	err := r.db.Select(&users, query)
	return users, err
}

func (r *UserPostgres) GetByUsername(username string) (blog.User, error) {
	var user blog.User
	query := fmt.Sprintf(`SELECT username, name FROM %s WHERE username=$1`, usersTable)
	err := r.db.Get(&user, query, username)
	return user, err
}