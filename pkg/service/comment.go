package service

import (
	"github.com/SergeiMurashev/blog-app/pkg/models"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
)

type CommentService struct {
	repo repository.Comment
}

func (s *CommentService) DeleteComment(comment models.CommentInputDelete) error {
	return s.repo.DeleteComment(comment)
}

func (s *CommentService) UpdateComment(comment models.CommentInputUpdate) (*models.Comment, error) {
	return s.repo.UpdateComment(comment)
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(comment models.CommentInputCreate) (*models.Comment, error) {
	return s.repo.CreateComment(comment)
}
