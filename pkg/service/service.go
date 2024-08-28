package service

import (
	models2 "github.com/SergeiMurashev/blog-app/models"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
)

// Post определяет методы, которые должны быть реализованы для работы с постами
type Post interface {
	CreatePost(post models2.PostInputCreate) (*models2.Post, error)
	DeletePost(post models2.PostInputDelete, email string) error
	UpdatePost(post models2.PostInputUpdate, email string) (*models2.Post, error)
	UserAuthorPost(email string, postID int) (bool, error)
}

// User определяет методы, которые должны быть реализованы для работы с пользователями
type User interface {
	CreateUser(user models2.UserInputCreate) (*models2.User, error)
	Authorization(email, password string) (*models2.AuthorizationOutput, error)
	ParseToken(token string) (string, error)
}

// Comment определяет методы, которые должны быть реализованы для работы с комментариями
type Comment interface {
	CreateComment(comment models2.CommentInputCreate) (*models2.Comment, error)
	DeleteComment(comment models2.CommentInputDelete, email string) error
	UpdateComment(comment models2.CommentInputUpdate, email string) (*models2.Comment, error)
	UserAuthorComment(email string, commentID int) (bool, error)
}

// Service объединяет все интерфейсы в один, представляя методы для работы с пользователями, постами и комментариями
type Service struct {
	User
	Post
	Comment
}

// NewService Создает новый объект Service, инициализируя его конкретными реализациями интерфейсов
func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repos.User),
		Post:    NewPostService(repos.Post),
		Comment: NewCommentService(repos.Comment),
	}
}
