package main

import (
	"log"
	"todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := &todo.Server{}
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("an error occurred while running the server:\n\t%s", err.Error())
	}
}
