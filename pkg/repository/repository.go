package repository

import (
	"github.com/SergeiMurashev/blog-app/pkg/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user models.UserInputCreate) (*models.User, error)
	GetUser(email string) (*models.User, error)
}

type Post interface {
	CreatePost(post models.PostInputCreate) (*models.Post, error)
	DeletePost(post models.PostInputDelete) error
	UpdatePost(post models.PostInputUpdate) (*models.Post, error)
	UserAuthorPost(email string, postID int) (bool, error)
}
type Comment interface {
	CreateComment(comment models.CommentInputCreate) (*models.Comment, error)
	DeleteComment(comment models.CommentInputDelete) error
	UpdateComment(comment models.CommentInputUpdate) (*models.Comment, error)
	UserAuthorComment(email string, commentID int) (bool, error)
}

type Repository struct {
	User
	Post
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Post:    NewPostPostgres(db),
		Comment: NewCommentPostgres(db),
	}
}
