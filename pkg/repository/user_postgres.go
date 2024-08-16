package repository

import (
	"fmt"
	"github.com/SergeiMurashev/blog-app/pkg/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}

}

func (r *UserPostgres) CreateUser(user models.UserInputCreate) (*models.User, error) {
	var output models.User
	err := r.db.Get(&output, `INSERT INTO "Users" ( name, email, password,create_date) values ($1, $2, $3, $4) RETURNING name, email, password, create_date`,
		user.Name,
		user.Email,
		user.Password,
		user.Create_date)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return &output, err
}

// ошибочка
func (r *UserPostgres) GetUser(username, password string) (*models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
