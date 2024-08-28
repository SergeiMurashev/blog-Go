package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Структура для ответа об ошибке
type errorResponse struct {
	Message string `json:"message"`
}

// Структура для ответа о статусе
type statusResponse struct {
	Status string `json:"status"`
}

// Функция для создания и отправки ответа об ошибке
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
