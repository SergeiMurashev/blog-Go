package models

import "time"

type User struct {
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	Create_date time.Time `json:"create_date" db:"create_date"`
}

type UserInputCreate struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"email"`
	Password    string `json:"password" binding:"gte=10"`
	Create_date string `json:"create_date"`
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
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Create_date time.Time `json:"create_date" db:"create_date"`
}
