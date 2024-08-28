package repository

import (
	"github.com/SergeiMurashev/blog-app/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// UserPostgres представляет работу с пользователями в базе данных Postgres
type UserPostgres struct {
	db *sqlx.DB
}

// Создание нового экземпляра UserPostgres
func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}

}

// CreateUser - добавляет нового пользователя в БД, user объект с данными юзера.
func (r *UserPostgres) CreateUser(user models.UserInputCreate) (*models.User, error) {
	var output models.User
	// Выполняем SQL-запрос для вставки нового пользователя и получаем его данные.
	err := r.db.Get(&output, `INSERT INTO "Users" ( name, email, password,create_date) values ($1, $2, $3, $4) RETURNING name, email, password, create_date`,
		user.Name,
		user.Email,
		user.Password,
		user.CreateDate)
	if err != nil {
		// Если произошла ошибка, записываем ее в Log
		logrus.Error(err.Error())
		return nil, err
	}
	return &output, err
}

// GetUser получает информацию о юзере по его email'у. Email - эл/почта изера.
func (r *UserPostgres) GetUser(email string) (*models.User, error) {
	var user models.User
	// Выполняем SQL-запрос для получения данных пользователя по его email
	err := r.db.Get(&user, `SELECT * FROM "Users" WHERE email=$1`, email)

	return &user, err
}
