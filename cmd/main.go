package main

import (
	"github.com/mukolla/go-todo-app"
	"github.com/mukolla/go-todo-app/pkg/handler"
	"github.com/mukolla/go-todo-app/pkg/repository"
	"github.com/mukolla/go-todo-app/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Must be start word: Bearer ... JWT authorization token

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	//logrus.SetLevel(logrus.DebugLevel)

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		logrus.Fatalf("failed to initialize DB %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
