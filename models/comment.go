package models

import "time"

type Comment struct {
	Id         int       `json:"id" db:"id"`
	Text       string    `json:"text" db:"text"`
	CreateDate time.Time `json:"createDate" db:"createDate"`
	Author     string    `json:"author" db:"author"`
	Post       int       `json:"post" db:"post"`
}

type CommentInputCreate struct {
	Text       string `json:"text" binding:"required"`
	CreateDate string `json:"createDate"`
	Author     string `json:"author"`
	Post       int    `json:"post"`
}

type CommentInputUpdate struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type CommentInputDelete struct {
	Id int `json:"id"`
}

type CommentInputCheck struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
}
