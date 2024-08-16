package repository

import (
	"github.com/SergeiMurashev/blog-app/pkg/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CommentPostgres struct {
	db *sqlx.DB
}

func NewCommentPostgres(db *sqlx.DB) *CommentPostgres {
	return &CommentPostgres{db: db}
}

func (r *CommentPostgres) CreateComment(comment models.CommentInputCreate) (*models.Comment, error) {
	var output models.Comment
	err := r.db.Get(&output, `INSERT INTO "Comment" ( text, create_date, author, post) values ($1, $2, $3, $4) RETURNING id,text, create_date, author, post`,
		comment.Text,
		comment.Create_date,
		comment.Author,
		comment.Post)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return &output, err
}

func (r *CommentPostgres) DeleteComment(comment models.CommentInputDelete) error {
	_, err := r.db.Exec(`DELETE FROM "Comment" WHERE id = $1`, comment.Id)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	return nil
}

func (r *CommentPostgres) UpdateComment(comment models.CommentInputUpdate) (*models.Comment, error) {
	var output models.Comment
	err := r.db.Get(&output, `UPDATE "Comment" SET text = $1 WHERE id = $2 RETURNING *`,
		comment.Text,
		comment.Id)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return &output, nil
}
