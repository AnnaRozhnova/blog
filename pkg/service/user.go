package service

import (
	"github.com/AnnaRozhnova/blog"
	"github.com/AnnaRozhnova/blog/pkg/repository"
)


type UserService struct {
	repo repository.User
}

func newUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}


func (s *UserService) GetAll() ([]blog.User, error) {
	return s.repo.GeyAll()
}