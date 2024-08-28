package handler

import (
	"github.com/SergeiMurashev/blog-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Обработчик для регистраций нового пользователя
func (h *Handler) signUp(c *gin.Context) {
	// Создаем переменную для хранения данных о новом пользователе
	var input models.UserInputCreate

	// Извлекаем данные из тела запроса и заполняем переменную input
	// Если данные некорректны, возвращаем ошибку
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	// Создаем пользователя, вызывая метод CreateUser из сервиса
	user, err := h.services.User.CreateUser(input)
	if err != nil {
		// Если при создании пользователя произошла ошибка, возвращаем ошибку
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// если все прошло успешно, возвращаем ID созданного пользователя
	c.JSON(http.StatusOK, gin.H{
		"id": user,
	})
}

// Обработчик для входа пользователя
func (h *Handler) signIn(c *gin.Context) {
	// Создание переменной для хранения данных для входа
	var input models.SignInInput

	// Извлекается данные из тела запроса и заполнением переменной input
	// Если данные некорректны, возвращаем ошибку
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// пытаемся авторизовать пользователя, вызывая метод Authorization из сервиса
	token, err := h.services.User.Authorization(input.Email, input.Password)
	if err != nil {
		// Если при авторизации произошла ошибка, возвращаем ошибку
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Если авторизация прошла успешно, возвращаем токен
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
