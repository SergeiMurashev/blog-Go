package blog

import (
	"context"
	"net/http"
	"time"
)

// Структура Server представляет HTTP-сервер
type Server struct {
	httpServer *http.Server
}

// Метод Run запускает HTTP-сервер на указанном порту с переданным обработчиком запросов
func (s *Server) Run(port string, handler http.Handler) error {
	// Настраиваем параметры сервера
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	// Запускаем сервер и начинаем принимать входящие HTTP-запросы
	return s.httpServer.ListenAndServe()
}

// Метод Shutdown останавливает сервер с возможностью плавного завершения работы
func (s *Server) Shutdown(ctx context.Context) error {

	return s.httpServer.Shutdown(ctx)
}
