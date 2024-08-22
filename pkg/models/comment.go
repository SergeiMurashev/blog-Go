package models

import "time"

type Comment struct {
	Id          int       `json:"id" db:"id"`
	Text        string    `json:"text" db:"text"`
	Create_date time.Time `json:"create_Date" db:"create_date"`
	Author      string    `json:"author" db:"author"`
	Post        int       `json:"post" db:"post"`
}

type CommentInputCreate struct {
	Text        string `json:"text" binding:"required"`
	Create_date string `json:"create_date"`
	Author      string `json:"author"`
	Post        int    `json:"post"`
}

type CommentInputUpdate struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type CommentInputDelete struct {
	Id int `json:"id"`
}
