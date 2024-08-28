package models

import "time"

type Post struct {
	Id         int       `json:"id" db:"id"`
	Title      string    `json:"title" db:"title"`
	Text       string    `json:"text" db:"text"`
	CreateDate time.Time `json:"createDate" db:"createDate"`
	Author     string    `json:"author" db:"author"`
}

type PostInputCreate struct {
	Title      string `json:"title" binding:"required"`
	Text       string `json:"text"  binding:"required"`
	CreateDate string `json:"createDate"`
	Author     string `json:"author"`
}

type PostInputUpdate struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type PostInputDelete struct {
	Id int `json:"id"`
}
type PostInputCheck struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
}
