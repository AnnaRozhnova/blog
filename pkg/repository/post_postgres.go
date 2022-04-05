package repository

import (
	"fmt"

	"github.com/AnnaRozhnova/blog"
	"github.com/jmoiron/sqlx"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) Create(post blog.Post) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, body, username) values ($1, $2, $3) RETURNING id", postsTable)
	row := r.db.QueryRow(query, post.Title, post.Body, post.Username)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PostPostgres) GetAll() ([]blog.Post, error) {
	var posts []blog.Post
	query := fmt.Sprintf("SELECT * FROM %s", postsTable)
	err := r.db.Select(&posts, query)
	return posts, err
}
func (r *PostPostgres) GetById(id int) (blog.Post, error) {
	var post blog.Post
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postsTable)
	err := r.db.Get(&post, query, id)
	return post, err
}


func (r *PostPostgres) GetByUsername(username string) ([]blog.Post, error) {
	var posts []blog.Post 
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = $1", postsTable)
	err := r.db.Select(&posts, query, username)
	return posts, err
} 