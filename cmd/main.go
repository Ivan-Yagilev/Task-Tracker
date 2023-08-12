package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("config init error: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("env reading error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("PG_HOST"),
		Port:     "5432",
		Username: os.Getenv("PG_USER"),
		DBName:   os.Getenv("PG_NAME"),
		SSLMode:  "disable",
		Password: os.Getenv("PG_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to init db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := &todo.Server{}
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("an error occurred while running the server: %s", err.Error())
		}
	}()

	logrus.Print("Application started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Application finished")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("server shutting down error: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("db connection close error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
