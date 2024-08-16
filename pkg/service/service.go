package service

import (
	"github.com/SergeiMurashev/blog-app/pkg/models"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
)

type Post interface {
	CreatePost(post models.PostInputCreate) (*models.Post, error)
	DeletePost(post models.PostInputDelete) error
	UpdatePost(post models.PostInputUpdate) (*models.Post, error)
}

type User interface {
	CreateUser(user models.UserInputCreate) (*models.User, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Comment interface {
	CreateComment(comment models.CommentInputCreate) (*models.Comment, error)
	DeleteComment(comment models.CommentInputDelete) error
	UpdateComment(comment models.CommentInputUpdate) (*models.Comment, error)
}

type Service struct {
	User
	Post
	Comment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repos.User),
		Post:    NewPostService(repos.Post),
		Comment: NewCommentService(repos.Comment),
	}
}
