package service

import (
	"errors"
	"github.com/SergeiMurashev/blog-app/pkg/models"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
	"github.com/sirupsen/logrus"
)

type PostService struct {
	repo repository.Post
}

func (s *PostService) UpdatePost(post models.PostInputUpdate, email string) (*models.Post, error) {
	// Является ли юзер автором поста
	exest, err := s.repo.UserAuthorPost(email, post.Id)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	// Если да, разрешаем апдейт
	if exest {
		return s.repo.UpdatePost(post)
	} else {
		return nil, errors.New("no your post")
	}
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post models.PostInputCreate) (*models.Post, error) {
	return s.repo.CreatePost(post)
}

func (s *PostService) DeletePost(post models.PostInputDelete, email string) error {
	//Является ли юзер автором поста
	exest, err := s.repo.UserAuthorPost(email, post.Id)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	// Если да, разрешаем удаление
	if exest {
		return s.repo.DeletePost(post)
	} else {
		return errors.New("no your post")
	}
}

func (s *PostService) UserAuthorPost(email string, postID int) (bool, error) {
	return s.repo.UserAuthorPost(email, postID)
}
