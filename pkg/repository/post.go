package repository

import (
	"github.com/SergeiMurashev/blog-app/pkg/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PostPostgres struct {
	db *sqlx.DB
}

func (r *PostPostgres) DeletePost(post models.PostInputDelete) error {
	_, err := r.db.Exec(`DELETE FROM "Post" WHERE id = $1`, post.Id)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	return nil
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}

}

func (r *PostPostgres) CreatePost(post models.PostInputCreate) (*models.Post, error) {
	var output models.Post
	err := r.db.Get(&output, `INSERT INTO "Post" ( title, text, create_date, author) values ($1, $2, $3, $4) RETURNING id, title, text, create_date, author`,
		post.Title,
		post.Text,
		post.Create_date,
		post.Author)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return &output, err
}

func (r *PostPostgres) UpdatePost(post models.PostInputUpdate) (*models.Post, error) {
	var output models.Post
	err := r.db.Get(&output, `UPDATE "Post" SET title = $1, text = $2 WHERE id = $3 RETURNING *`,
		post.Title,
		post.Text,
		post.Id)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return &output, nil
}

func (r *PostPostgres) UserAuthorPost(email string, postID int) (bool, error) {
	var exist bool
	err := r.db.Get(&exist, `SELECT EXISTS(select * FROM "Post" WHERE id = $1, AND public."Post".author = $2)`, postID, email)
	if err != nil {
		logrus.Error(err.Error())
		return false, err
	}
	return exist, nil
}
