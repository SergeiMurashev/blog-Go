package repository

import (
	models2 "github.com/SergeiMurashev/blog-app/models"
	"github.com/jmoiron/sqlx"
)

// User определяте методы для работы с пользователями в БД
type User interface {
	CreateUser(user models2.UserInputCreate) (*models2.User, error)
	GetUser(email string) (*models2.User, error)
}

// Post определяет методы для работы с постами в БД
type Post interface {
	CreatePost(post models2.PostInputCreate) (*models2.Post, error)
	DeletePost(post models2.PostInputDelete) error
	UpdatePost(post models2.PostInputUpdate) (*models2.Post, error)
	UserAuthorPost(email string, postID int) (bool, error)
}

// Comment определяет методы для работы с комментариями в БД
type Comment interface {
	CreateComment(comment models2.CommentInputCreate) (*models2.Comment, error)
	DeleteComment(comment models2.CommentInputDelete) error
	UpdateComment(comment models2.CommentInputUpdate) (*models2.Comment, error)
	UserAuthorComment(email string, commentID int) (bool, error)
}

// Объединяет все интерфейсы в один, это позволяет работать с user, post, comment через один объект
type Repository struct {
	User
	Post
	Comment
}

// Создается новый объект Repository db - подключение к БД
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Post:    NewPostPostgres(db),
		Comment: NewCommentPostgres(db),
	}
}
