package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	accountContextName  = "account-context"
	accountTokenName    = "account-token"
	authorizationHeader = "Authorization"
)

// Проверка и извлечение информации о пользователе из заголовка авторизаций
func (h *Handler) userIdentity(c *gin.Context) {
	// Получение значение заголовка авторизация из запроса
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		// Если заголовок пуст, возвращаем ошибку 401 (не авторизован)
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	// Пытаемся извлечь email пользователя из токена
	email, err := h.services.User.ParseToken(header)
	if err != nil {
		// Если произошла ошибка при разборе токена, возвращаем ошибку 401 (не авторизован)
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Сохраняем email токен в контекст запроса
	c.Set(accountContextName, email)
	c.Set(accountTokenName, header)
}

// Функция для получения email  пользователя из заголовка авторизаций
func (h *Handler) GetUser(c *gin.Context) string {
	// Получаем значения заголовка авторизации из запроса
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		// Если заголовок пуст, возвращаем ошибку 401 (не авторизован)
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return ""
	}
	// Пытаемся извлечь email пользователя из токена
	email, err := h.services.User.ParseToken(header)
	if err != nil {
		// Если произошла ошибка при разборе токена, возвращаем ошибку 401 (не авторизован)
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return ""
	}
	// Возвращаем email пользователя
	return email
}
