package repository

import (
	"github.com/SergeiMurashev/blog-app/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// PostPostgres управляет операциями с постами в БД
type PostPostgres struct {
	db *sqlx.DB
}

// Удаление поста из БД по его идинтификатору. Post - объект с данными, включая id для удаления
func (r *PostPostgres) DeletePost(post models.PostInputDelete) error {
	_, err := r.db.Exec(`DELETE FROM "Post" WHERE id = $1`, post.Id)
	if err != nil {
		// Если ошибка при удалений, запись в Log.
		logrus.Error(err.Error())
		return err
	}
	// Возвращаем nil, если все прошло успешно.
	return nil
}

// Создание нового экземпляра PostPostgres.
func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}

}

// Добавление нового поста в БД, Post: объект с данными нового поста.
func (r *PostPostgres) CreatePost(post models.PostInputCreate) (*models.Post, error) {
	// Объявляет переменную для хранения данных о посте, которая будет заполнена результатами SQL-запроса.
	var output models.Post
	// Выполняется сам SQL-запрос для вставки нового поста и получения его данных
	err := r.db.Get(&output, `INSERT INTO "Post" ( title, text, create_date, author) values ($1, $2, $3, $4) RETURNING id, title, text, create_date, author`,
		post.Title,
		post.Text,
		post.CreateDate,
		post.Author)
	if err != nil {
		// Если ошибка при создании поста, записываем ее в Log и возвращаем ошибку.
		logrus.Error(err.Error())
		return nil, err
	}
	// Возвращаем созданный пост и nil как ошибку
	return &output, err
}

// Обновляем данные поста в БД, post: - объект с новыми данными и идентификатором Id поста для обновления
func (r *PostPostgres) UpdatePost(post models.PostInputUpdate) (*models.Post, error) {
	// Объявляет переменную для хранения данных о посте, которая будет заполнена результатами SQL-запроса.
	var output models.Post
	// Выполняем SQL-запрос для обновления данных поста и получаем его обновленные данные
	err := r.db.Get(&output, `UPDATE "Post" SET title = $1, text = $2 WHERE id = $3 RETURNING *`,
		post.Title,
		post.Text,
		post.Id)
	if err != nil {
		// Если ошибка при обновлении поста, записываем ее в Log и возвращаем ошибку.
		logrus.Error(err.Error())
		return nil, err
	}
	// Возвращаем обновленный пост и nil как ошибку
	return &output, nil
}

// UserAuthorPost проверяет ,является ли пользователь автором поста
func (r *PostPostgres) UserAuthorPost(email string, postID int) (bool, error) {
	var exist bool
	// Выполняем SQL-запрос, чтобы проверить, существует ли пост с указанным идентификатором (id) и принадлежит ли он пользователю
	err := r.db.Get(&exist, `SELECT EXISTS(select * FROM "Post" WHERE id = $1, AND public."Post".author = $2)`, postID, email)
	if err != nil {
		// Если возникла ошибка при проверке, записываем её в log и возвращаем false
		logrus.Error(err.Error())
		return false, err
	}
	// Возвращаем результат проверки и nil как ошибку
	return exist, nil
}
