package service

import (
	"errors"
	"github.com/SergeiMurashev/blog-app/models"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
	"github.com/sirupsen/logrus"
)

// PostService предоставляет методы для работы с постами
type PostService struct {
	repo repository.Post
}

// метод для обновления комментария - UpdatePost, для струткуры - PostService.
// аргументы - post объект типа models.PostInputUpdate - содержит данные для обновления, email - юзер чей пост.
func (s *PostService) UpdatePost(post models.PostInputUpdate, email string) (*models.Post, error) {
	// Проверка существования поста. Вызывается метод s.repo, проверяет сущ ли пост с указанным id и принадлежит ли он юзеру с данным email
	// Exest (от слова exists - существовать)- переменная, bool'евого типа, то есть true если коммент с id есть и его email, то правда. Если юзер не он коммента, то false.
	// Для понятности можно вместо exest написать isAuthor.
	isExist, err := s.repo.UserAuthorPost(email, post.Id)
	if err != nil {
		// Если возникла ошибка при проверке, записываем ее в лог и возвращаем ошибку
		logrus.Error(err.Error())
		return nil, err
	}
	// если коммент принадлежит пользователю, обновляем его
	if isExist {
		return s.repo.UpdatePost(post)
	} else {
		// Если комментарий не его, возвращаем ошибку.
		return nil, errors.New("no your post")
	}
}

// Создает новый экземпляр PostService, repo: объект для работы с постами в БД
func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

// Создает новый пост и сохраняет его в БД
func (s *PostService) CreatePost(post models.PostInputCreate) (*models.Post, error) {
	return s.repo.CreatePost(post)
}

// метод для обновления комментария - DeletePost, для струткуры - PostService.
// аргументы - post объект типа models.PostInputDelete - содержит данные для удаления, email - юзер чей коммент.
func (s *PostService) DeletePost(post models.PostInputDelete, email string) error {
	// Проверка существования поста. Вызывается метод s.repo, проверяет сущ ли пост с указанным id и принадлежит ли он юзеру с данным email
	// Exest (от слова exists - существовать)- переменная, bool'евого типа, то есть true если коммент с id есть и его email, то правда. Если юзер не он коммента, то false.
	// Для понятности можно вместо exest написать isAuthor.
	isAuthor, err := s.repo.UserAuthorPost(email, post.Id)
	if err != nil {
		// Если возникла ошибка при проверке, записываем ее в лог и возвращаем ошибку
		// Log - диагностика и отладка (позволяют понять что происходит внутри)
		logrus.Error(err.Error())
		return err
	}
	// если пост принадлежит пользователю, удаляем его.
	if isAuthor {
		return s.repo.DeletePost(post)
	} else {
		// если нет, то возвращаем ошибку.
		return errors.New("no your post")
	}
}

// Проверяет принадлежит ли пост пользователю, принимает emfil и postID возвращает bool'евое значение true или false или ошибку, если произошла проблема
func (s *PostService) UserAuthorPost(email string, postID int) (bool, error) {
	return s.repo.UserAuthorPost(email, postID)
}
