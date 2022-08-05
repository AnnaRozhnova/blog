package service

import (
	"github.com/AnnaRozhnova/blog"
	"github.com/AnnaRozhnova/blog/pkg/repository"
)

type Authorization interface {
	CreateUser(user blog.User) error
	GetUser(username, password string) (blog.User, error)
}

type User interface {
	GetAll() ([]blog.User, error)
	GetByUsername(username string) (blog.User, error)
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

type Service struct {
	Authorization
	User
	Post
	Comment
}


func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		User: 		   newUserService(repo.User),		
		Post:          NewPostService(repo.Post),
		Comment:	   NewCommentService(repo.Comment),
	}
}
