package service

import (
	"errors"
	"github.com/SergeiMurashev/blog-app/pkg/models"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
	"github.com/sirupsen/logrus"
)

type CommentService struct {
	repo repository.Comment
}

func (s *CommentService) DeleteComment(comment models.CommentInputDelete, email string) error {
	// Является ли юзер автором коммента
	isAuthor, err := s.repo.UserAuthorComment(email, comment.Id)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	// Если да, разрешаем удаление
	if isAuthor {
		return s.repo.DeleteComment(comment)
	} else {
		return errors.New("no your comment")
	}
}

func (s *CommentService) UpdateComment(comment models.CommentInputUpdate, email string) (*models.Comment, error) {
	// Является ли юзер автором коммента
	exest, err := s.repo.UserAuthorComment(email, comment.Id)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	// Если да, разрешаем апдейт
	if exest {
		return s.repo.UpdateComment(comment)
	} else {
		return nil, errors.New("no your comment")
	}
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(comment models.CommentInputCreate) (*models.Comment, error) {
	return s.repo.CreateComment(comment)
}
func (s *CommentService) UserAuthorComment(email string, commentID int) (bool, error) {
	return s.repo.UserAuthorComment(email, commentID)
}
