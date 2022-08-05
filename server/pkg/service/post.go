package service

import (
	"github.com/AnnaRozhnova/blog"
	"github.com/AnnaRozhnova/blog/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) Create(post blog.Post) (int, error) {
	return s.repo.Create(post)
}

func (s *PostService) GetAll() ([]blog.Post, error) {
	return s.repo.GetAll()
}
func (s *PostService) GetById(id int) (blog.Post, error) {
	return s.repo.GetById(id)
}

func (s *PostService) GetByUsername(username string) ([]blog.Post, error) {
	return s.repo.GetByUsername(username)
}
