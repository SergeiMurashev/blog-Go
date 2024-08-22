package service

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"github.com/SergeiMurashev/blog-app/pkg/models"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"sync"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var session = sync.Map{}

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.UserInputCreate) (*models.User, error) {
	user.Password = getPasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func getPasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error(err.Error())
		return ""
	}

	return string(hash)
}

func (s *UserService) Authorization(email, password string) (*models.AuthorizationOutput, error) {
	account, err := s.repo.GetUser(email)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	err = comparePasswords(account.Password, password)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	token := s.GenerateToken(email)
	var accountOutput = models.UserOutput{
		Name:        account.Name,
		Email:       account.Email,
		Create_date: account.Create_date,
	}
	var output = models.AuthorizationOutput{
		User:  accountOutput,
		Token: token,
	}
	return &output, nil
}

func comparePasswords(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (s *UserService) GenerateToken(email string) string {
	token := generateToken()
	session.Store(token, email)
	return token
}

func (s *UserService) ParseToken(accessToken string) (string, error) {
	email, ok := session.Load(accessToken)
	if !ok {
		logrus.Warnf("err auth")
		return "", errors.New("err auth")
	}
	return email.(string), nil
}

func generateToken() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 1024)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	s := sha512.New()
	s.Write([]byte(string(b)))
	return hex.EncodeToString(s.Sum(nil))
}
