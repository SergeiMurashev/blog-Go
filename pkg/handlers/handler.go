package handler

import (
	"github.com/SergeiMurashev/blog-app/pkg/service"
	"github.com/gin-gonic/gin"
)

// Ссылка на структуру с сервисами для выполнения операций
type Handler struct {
	services *service.Service
}

// Конструктор для создания нового обработчика
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// Метод для инициализации маршрутов (роутеров) приложения
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New() // Создаем новый экземпляр маршрутизатора Gin
	// Группируем маршруты под "/api"
	api := router.Group("/api")
	{
		// Группируем маршруты для работы с пользователями под "/api/user"
		user := api.Group("/user")
		{
			user.POST("/sign-up", h.signUp)
			user.POST("/sign-in", h.signIn)
		}
		// Группируем маршруты для работы с постами под "/api/posts" и добавляем middleware для проверки авторизации
		Posts := api.Group("/posts", h.userIdentity)
		{
			Posts.POST("", h.createPost)
			Posts.PUT("/:id", h.updatePost)
			Posts.DELETE("/:id", h.deletePost)
		}
		// Группируем маршруты для работы с комментариями под "/api/comment"
		comment := api.Group("/comment")
		{
			comment.POST("/sign-up", h.createComment)
			comment.PUT("/:id", h.updateComment)
			comment.DELETE("/:id", h.deleteComment)
		}
	}
	// Возвращаем настроенный маршрутизатор
	return router
}
