package models

import "time"

type User struct {
	Name       string    `json:"name" db:"name"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"password" db:"password"`
	CreateDate time.Time `json:"createDate" db:"createDate"`
}

type UserInputCreate struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"email"`
	Password   string `json:"password" binding:"gte=10"`
	CreateDate string `json:"createDate"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthorizationOutput struct {
	User  UserOutput `json:"user"`
	Token string     `json:"token"`
}

type UserOutput struct {
	Name       string    `json:"name" db:"name"`
	Email      string    `json:"email" db:"email"`
	CreateDate time.Time `json:"createDate" db:"createDate"`
}
