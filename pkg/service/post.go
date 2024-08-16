package service

import (
	"github.com/SergeiMurashev/blog-app/pkg/models"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func (s *PostService) UpdatePost(post models.PostInputUpdate) (*models.Post, error) {
	return s.repo.UpdatePost(post)
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post models.PostInputCreate) (*models.Post, error) {
	return s.repo.CreatePost(post)
}

func (s *PostService) DeletePost(post models.PostInputDelete) error {
	return s.repo.DeletePost(post)
}
