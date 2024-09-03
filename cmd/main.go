package main

import (
	"encoding/json"
	"os"

	"github.com/SergeiMurashev/blog-app"
	"github.com/SergeiMurashev/blog-app/pkg/handlers"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
	"github.com/SergeiMurashev/blog-app/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Для запуска через контейнер docker-compose up --build blog-app
// Порт на локалке 5434 но если из контейнера ?

func main() {
	// Настраиваем формат логов в формате JSON
	logrus.SetFormatter(new(logrus.JSONFormatter))
	// Загружаем конфигурацию приложения из файла config.json
	config, err := getConfig()
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	// Загружаем переменные окружения из файла .env
	if err := godotenv.Load(); err != nil {
		// Если не удалось загрузить переменные окружения, выводим сообщение и завершаем работу
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	// Подключаемся к базе данных Postgres с использованием загруженной конфигурации
	db, err := repository.NewPostgresDB(config)
	if err != nil {
		// Если не удалось подключиться к базе данных, выводим сообщение и завершаем работу
		logrus.Fatalf("failed to initialize database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	// Создаем новый сервер и запускаем его на порту 25504
	srv := new(blog.Server)
	if err := srv.Run("25504", handlers.InitRoutes()); err != nil {
		// Если возникла ошибка при запуске сервера, выводим сообщение и завершаем работу
		logrus.Fatalf("error occurred while running HTTP server: %s", err.Error())
	}
}

// Функция для загрузки конфигурации из файла config.json
func getConfig() (repository.Config, error) {
	var config repository.Config
	// Читаем содержимое файла config.json
	raw, err := os.ReadFile("./configs/config.json")
	if err != nil {
		return config, err // Если не удалось прочитать файл, возвращаем ошибку
	}
	// Декодируем JSON в структуру конфигурации
	err = json.Unmarshal(raw, &config)
	if err != nil {
		return config, err // Если не удалось декодировать JSON, возвращаем ошибку
	}

	return config, nil // Возвращаем загруженную конфигурацию
}
