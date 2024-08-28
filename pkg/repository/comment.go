package repository

import (
	"github.com/SergeiMurashev/blog-app/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// CommentPostgres управляет операциями с постами в БД
type CommentPostgres struct {
	db *sqlx.DB
}

// Создание нового экземпляра CommentPostgres, подключене db к БД
func NewCommentPostgres(db *sqlx.DB) *CommentPostgres {
	return &CommentPostgres{db: db}
}

// Добавление нового комментария в БД, Сomment: объект с данными нового комментария.
func (r *CommentPostgres) CreateComment(comment models.CommentInputCreate) (*models.Comment, error) {
	// Объявляет переменную для хранения данных о комментарий, которая будет заполнена результатами SQL-запроса.
	var output models.Comment
	// Выполняется сам SQL-запрос для вставки нового комментария и получения его данных
	err := r.db.Get(&output, `INSERT INTO "Comment" ( text, create_date, author, post) values ($1, $2, $3, $4) RETURNING id,text, create_date, author, post`,
		comment.Text,
		comment.CreateDate,
		comment.Author,
		comment.Post)
	if err != nil {
		// Если ошибка при создании комментария, записываем ее в Log и возвращаем ошибку.
		logrus.Error(err.Error())
		return nil, err
	}
	// Возвращаем созданный комментария и nil как ошибку
	return &output, err
}

// Удаление комментария из БД по его идинтификатору (id). Comment - объект с данными, включая id для удаления
func (r *CommentPostgres) DeleteComment(comment models.CommentInputDelete) error {
	_, err := r.db.Exec(`DELETE FROM "Comment" WHERE id = $1`, comment.Id)
	if err != nil {
		// Если ошибка при удалений, запись в log
		logrus.Error(err.Error())
		return err
	}
	// Возвращаем nil если все прошло успешно
	return nil
}

// Обновляем данные комментария в БД, comment: - объект с новыми данными и идентификатором Id комментария для обновления
func (r *CommentPostgres) UpdateComment(comment models.CommentInputUpdate) (*models.Comment, error) {
	// Объявляет переменную для хранения данных о комментарии, которая будет заполнена результатами SQL-запроса.
	var output models.Comment
	// Выполняем SQL-запрос для обновления данных комментария и получаем его обновленные данные
	err := r.db.Get(&output, `UPDATE "Comment" SET text = $1 WHERE id = $2 RETURNING *`,
		comment.Text,
		comment.Id)
	if err != nil {
		// Если ошибка при обновлении комментария, записываем ее в Log и возвращаем ошибку.
		logrus.Error(err.Error())
		return nil, err
	}
	// Возвращаем обновленный комментария и nil как ошибку
	return &output, nil
}

// UserAuthorComment проверяет ,является ли пользователь автором комментария
func (r *CommentPostgres) UserAuthorComment(email string, commentID int) (bool, error) {
	var exist bool
	// Выполняем SQL-запрос, чтобы проверить, существует ли комментарий с указанным идентификатором (id) и принадлежит ли он пользователю
	query := `SELECT EXISTS(SELECT 1 FROM "Comment" WHERE id = $1 AND author = $2)`
	err := r.db.Get(&exist, query, commentID, email)
	if err != nil {
		// Если возникла ошибка при проверке, записываем её в log и возвращаем false
		logrus.WithError(err).Error("Error checking if user is the author of the comment")
		return false, err
	}
	// Возвращаем результат проверки и nol как ошибку
	return exist, nil
}
