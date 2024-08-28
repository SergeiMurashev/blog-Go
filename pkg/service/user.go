package service

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"github.com/SergeiMurashev/blog-app/models"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"sync"
	"time"
)

// Набор символов для генерации токенов
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// Хранилище активных сессий
var session = sync.Map{}

// Методы для работы с пользователями
type UserService struct {
	repo repository.User
}

// Новый экземпляр UserService
func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

// CreateUser создает нового пользователя и сохраняет его в репозитории
func (s *UserService) CreateUser(user models.UserInputCreate) (*models.User, error) {
	// Хешируем пароль пользователя перед сохранением
	user.Password = getPasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// getPasswordHash хеширует пароль с использованием bcrypt
func getPasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// Записываем ошибку в log, если хэш неудался.
		logrus.Error(err.Error())
		return ""
	}

	return string(hash)
}

// Authorization проверяет учетные данные пользователя и выдает токен
func (s *UserService) Authorization(email, password string) (*models.AuthorizationOutput, error) {
	// Получаем данные пользователя по email
	account, err := s.repo.GetUser(email)
	if err != nil {
		// Записываем ошибку в log.
		logrus.Error(err.Error())
		return nil, err
	}
	// Сравниваем предоставленный пароль с хэшированным паролем
	err = comparePasswords(account.Password, password)
	if err != nil {
		// Записываем ошибку в log.
		logrus.Error(err.Error())
		return nil, err
	}
	// Генерируем токен для авторизированного пользователя
	token := s.GenerateToken(email)

	// Формируем выходные данные
	var accountOutput = models.UserOutput{
		Name:       account.Name,
		Email:      account.Email,
		CreateDate: account.CreateDate,
	}
	var output = models.AuthorizationOutput{
		User:  accountOutput,
		Token: token,
	}
	return &output, nil
}

// comparePasswords сравнивает хешированный пароль с предоставленным
func comparePasswords(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// GenerateToken создает новый токен и сохраняет его в хранилище сессий
func (s *UserService) GenerateToken(email string) string {
	token := generateToken()
	session.Store(token, email)
	return token
}

// ParseToken извлекает email из токена, если токен существует в хранилище сессий
func (s *UserService) ParseToken(accessToken string) (string, error) {
	email, ok := session.Load(accessToken)
	if !ok {
		// Записываем ошибки в log, если токен не найден.
		logrus.Warnf("err auth")
		return "", errors.New("err auth")
	}
	return email.(string), nil
}

// generateToken создает уникальный токен, используя случайные символы и хеширование
func generateToken() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 1024)
	// Записываем массив случайным символом
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	// Создаем хэш sha512 из случайных символов
	s := sha512.New()
	s.Write([]byte(string(b)))
	return hex.EncodeToString(s.Sum(nil))
}
