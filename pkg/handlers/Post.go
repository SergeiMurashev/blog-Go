package handler

import (
	"github.com/SergeiMurashev/blog-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Метод для создания нового поста
func (h *Handler) createPost(c *gin.Context) {
	// Создаем переменную для хранения данных о новом посте
	var input models.PostInputCreate

	// Извлекаем данные из тела запроса и заполняем переменную input
	// Если данные некорректны, возвращаем ошибку
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	// Получаем email пользователя, который делает запрос
	email := h.GetUser(c)
	// Устанавливаем автора поста
	input.Author = email
	// Создаем новый пост, вызывая метод CreatePost из сервиса
	post, err := h.services.Post.CreatePost(input)
	if err != nil {
		// Если возникла ошибка при создании поста, возвращаем ошибку
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Если пост создан успешно, возвращаем ID нового поста
	c.JSON(http.StatusOK, gin.H{
		"id": post,
	})
}

// Метод для обновления существующего поста
func (h *Handler) updatePost(c *gin.Context) {
	// Создаем переменную для хранения данных для обновления поста
	var input models.PostInputUpdate

	// Извлекаем ID поста из параметров URL и преобразуем его в число
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Если ID некорректно, возвращаем ошибку
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	input.Id = id

	// Извлекаем данные из тела запроса и заполняем переменную input
	// Если данные некорректны, возвращаем ошибку
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	// Получаем email пользователя, который делает запрос
	email := h.GetUser(c)

	// Обновляем пост, вызывая метод UpdatePost из сервиса
	output, err := h.services.Post.UpdatePost(input, email)
	if err != nil {
		// Если возникла ошибка при обновлении поста, возвращаем ошибку
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Если пост обновлен успешно, возвращаем обновленные данные поста
	c.JSON(http.StatusOK, output)
}

// Метод для удаления поста
func (h *Handler) deletePost(c *gin.Context) {
	// Создаем переменную для хранения данных о посте, который нужно удалить
	var input models.PostInputDelete

	// Извлекаем ID поста из параметров URL и преобразуем его в число
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Если ID некорректен, возвращаем ошибку
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	input.Id = id

	// Получаем email пользователя, который делает запрос
	email := h.GetUser(c)

	// Удаляем пост, вызывая метод DeletePost из сервиса
	err = h.services.Post.DeletePost(input, email)
	if err != nil {
		// Если возникла ошибка при удалений поста, возвращаем ошибку
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Если пост удален успешно, возвращаем ответ с статусом "ок" и кодом состояния 204 (No content)
	c.JSON(http.StatusNoContent, statusResponse{
		Status: "ok",
	})
}
