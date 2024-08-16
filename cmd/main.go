package main

import (
	"encoding/json"
	"github.com/SergeiMurashev/blog-app"
	"github.com/SergeiMurashev/blog-app/pkg/handlers"
	"github.com/SergeiMurashev/blog-app/pkg/repository"
	"github.com/SergeiMurashev/blog-app/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Config2 struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func main() {
	var config repository.Config
	logrus.SetFormatter(new(logrus.JSONFormatter))
	raw, err := os.ReadFile("./configs/config.json")
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	err = json.Unmarshal(raw, &config)
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(config)
	if err != nil {
		logrus.Fatalf(err.Error())
		return
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(blog.Server)
	if err := srv.Run("25504", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http srver: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
