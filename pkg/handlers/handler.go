package handler

import (
	"github.com/SergeiMurashev/blog-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/sign-up", h.signUp)
			user.POST("/sign-in", h.signIn)
		}

		Posts := api.Group("/posts")
		{
			Posts.POST("", h.createPost)
			Posts.PUT("/:id", h.updatePost)
			Posts.DELETE("/:id", h.deletePost)
		}
		comment := api.Group("/comment")
		{
			comment.POST("/sign-up", h.CreateComment)
			comment.PUT("/:id", h.updateComment)
			comment.DELETE("/:id", h.deleteComment)
		}
	}
	return router
}
