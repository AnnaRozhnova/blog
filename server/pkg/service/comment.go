package service

import (
	"github.com/AnnaRozhnova/blog"
	"github.com/AnnaRozhnova/blog/pkg/repository"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) Create(comment blog.Comment) error {
	return s.repo.Create(comment)
}

func (s *CommentService) GetByPostId(postId int) ([]blog.Comment, error) {
	return s.repo.GetByPostId(postId)
}
