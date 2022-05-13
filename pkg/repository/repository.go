package repository

import (
	"github.com/AnnaRozhnova/blog"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user blog.User) error
	GetUser(username, password string) (blog.User, error)
}
type User interface {
	GeyAll() ([]blog.User, error)
}
type Post interface {
	Create(post blog.Post) (int, error)
	GetAll() ([]blog.Post, error)
	GetById(id int) (blog.Post, error)
	GetByUsername(username string) ([]blog.Post, error)
}

type Comment interface {
	Create(comment blog.Comment) error
	GetByPostId(postId int) ([]blog.Comment, error)
}

type Repository struct {
	Authorization
	User
	Post
	Comment
}



func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User: 		   NewUserPostgres(db),
		Post:          NewPostPostgres(db),
		Comment: 	   NewCommentPostgres(db),
	}
}
