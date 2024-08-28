package service

import (
	"errors"
	"github.com/SergeiMurashev/blog-app/models"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
	"github.com/sirupsen/logrus"
)

// CommentService предоставляет методы для работы с комментариями
type CommentService struct {
	repo repository.Comment
}

// метод для обновления комментария - DeleteComment, для струткуры - CommentService.
// аргументы - comment объект типа models.CommentInputDelete - содержит данные для обновления, email - юзер чей коммент.
func (s *CommentService) DeleteComment(comment models.CommentInputDelete, email string) error {
	// Проверяем, принадлежит ли комментарий с данным Id пользователю с указанным email.
	isAuthor, err := s.repo.UserAuthorComment(email, comment.Id)
	if err != nil {
		// Если ошибка при проверке, записывам в log и возвращаем ошибку.
		logrus.Error(err.Error())
		return err
	}
	// Если комментарий его, то удаляем.
	if isAuthor {
		return s.repo.DeleteComment(comment)
	} else {
		// Если не его, возвращаем ошибку.
		return errors.New("no your comment")
	}
}

// Метод для обновления комментария - UpdateComment, для струткуры - CommentService.
// Аргументы - comment объект типа models.CommentInputUpdate - содержит данные для обновления, email - юзер чей коммент.
func (s *CommentService) UpdateComment(comment models.CommentInputUpdate, email string) (*models.Comment, error) {
	// Проверка существования коммента. Вызывается метод s.repo, проверяет сущ ли коммент с указанным id и принадлежит ли он юзеру с данным email
	// Exest (от слова exists - существовать)- переменная, bool'евого типа, то есть true если коммент с id есть и его email, то правда. Если юзер не от коммента, то false.
	// Для понятности можно вместо exest написать isAuthor.
	exest, err := s.repo.UserAuthorComment(email, comment.Id)
	if err != nil {
		// Если возникла ошибка при проверке, записываем ее в лог и возвращаем ошибку
		// Log - диагностика и отладка (позволяют понять что происходит внутри)
		logrus.Error(err.Error())
		return nil, err
	}
	// если коммент принадлежит пользователю, обновляем его.
	if exest {
		return s.repo.UpdateComment(comment)
	} else {
		// Если комментарий не его, возвращаем ошибку.
		return nil, errors.New("no your comment")
	}
}

// Создает новый экземпляр CommentService, repo: объект для работы с комментариями в БД
func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo: repo}
}

// Создает новый комментарий и сохраняет его в БД
func (s *CommentService) CreateComment(comment models.CommentInputCreate) (*models.Comment, error) {
	return s.repo.CreateComment(comment)
}

// Проверяет принадлежит ли комментарий пользователю, принимает emfil и commentID возвращает bool'евое значение true или false или ошибку, если произошла проблема
func (s *CommentService) UserAuthorComment(email string, commentID int) (bool, error) {
	return s.repo.UserAuthorComment(email, commentID)
}
