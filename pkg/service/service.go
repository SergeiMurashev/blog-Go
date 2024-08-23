package service

import (
	"github.com/SergeiMurashev/blog-app/pkg/models"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
)

type Post interface {
	CreatePost(post models.PostInputCreate) (*models.Post, error)
	DeletePost(post models.PostInputDelete, email string) error
	UpdatePost(post models.PostInputUpdate, email string) (*models.Post, error)
	UserAuthorPost(email string, postID int) (bool, error)
}

type User interface {
	CreateUser(user models.UserInputCreate) (*models.User, error)
	Authorization(email, password string) (*models.AuthorizationOutput, error)
	ParseToken(token string) (string, error)
}

type Comment interface {
	CreateComment(comment models.CommentInputCreate) (*models.Comment, error)
	DeleteComment(comment models.CommentInputDelete, email string) error
	UpdateComment(comment models.CommentInputUpdate, email string) (*models.Comment, error)
	UserAuthorComment(email string, commentID int) (bool, error)
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
