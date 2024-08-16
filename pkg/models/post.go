package models

import "time"

type Post struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Text        string    `json:"text" db:"text"`
	Create_date time.Time `json:"create_date" db:"create_date"`
	Author      string    `json:"author" db:"author"`
}

type PostInputCreate struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	Create_date string `json:"create_date"`
	Author      string `json:"author"`
}

type PostInputUpdate struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type PostInputDelete struct {
	Id int `json:"id"`
}
