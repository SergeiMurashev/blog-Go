package models

import "time"

type User struct {
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	Create_date time.Time `json:"create_date" db:"create_date"`
}

type UserInputCreate struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Create_date string `json:"create_date"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
