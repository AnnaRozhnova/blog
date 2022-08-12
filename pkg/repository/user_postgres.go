package repository

import (
	"fmt"

	"github.com/AnnaRozhnova/blog"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}



// NewUserPostgres creates new UserPostgres instance
func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}


// GeyAll gets all users from database
func (r *UserPostgres) GeyAll() ([]blog.User, error) {
	var users []blog.User
	query := fmt.Sprintf("SELECT username, name FROM %s", usersTable)
	err := r.db.Select(&users, query)
	return users, err
}

// GetByUsername get userinfo of particular user
func (r *UserPostgres) GetByUsername(username string) (blog.User, error) {
	var user blog.User
	query := fmt.Sprintf(`SELECT username, name FROM %s WHERE username=$1`, usersTable)
	err := r.db.Get(&user, query, username)
	return user, err
}