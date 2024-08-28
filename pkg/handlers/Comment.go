package handler

import (
	"github.com/SergeiMurashev/blog-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Метод для создания нового комментария
func (h *Handler) createComment(c *gin.Context) {
	// Создаем переменную для хранения данных о новом комментарии
	var input models.CommentInputCreate

	// Извлекаем данные из тела запроса и заполняем переменную input
	// Если данные некорректны, возвращаем ошибку с кодом 400 (Bad Request)
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	// Вызываем метод сервиса для создания комментария с полученными данными
	comment, err := h.services.Comment.CreateComment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Если создание прошло успешно, возвращаем ID нового комментария
	c.JSON(http.StatusOK, gin.H{
		"id": comment,
	})
}

// Метод для обновления существующего комментария
func (h *Handler) updateComment(c *gin.Context) {
	// Создаем переменную для хранения данных обновляемого комментария
	var input models.CommentInputUpdate
	// Получаем ID комментария из параметра URL преобразуем его в число
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	input.Id = id // Сохраняем ID в структуре input

	// Извлекаем данные из тела запроса и заполняем переменную input, если данные некорректны, возвращаем ошибку
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	// Получаем email автора комментария из токена, переданного в заголовке
	email := h.GetUser(c)
	// Обновляем комментарии с новыми данными
	output, err := h.services.Comment.UpdateComment(input, email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Если обновление прошло успешно, возвращаем обновленные данные комментария
	c.JSON(http.StatusOK, output)
}

// Метод для удаления комментария
func (h *Handler) deleteComment(c *gin.Context) {
	// Создание переменной для хранения данных удаляемого комментария
	var input models.CommentInputDelete
	// Получаем ID комментария из параметра URL и преобразуем его в число
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Если ID некорректный, возвращаем ошибку
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	input.Id = id // Сохраняем ID в структуре input
	// Получаем email автора комментария из токена, переданного в заголовке
	email := h.GetUser(c)
	// Удаляем комментарий
	err = h.services.Comment.DeleteComment(input, email)
	if err != nil {
		// Если возникла ошибка на сервере, возвращаем код 500
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Если удаление прошло успешно, возвращаем статус "Ok" и код 204 (No Content)
	c.JSON(http.StatusNoContent, statusResponse{
		Status: "ok",
	})
}
