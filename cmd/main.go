package main

import (
	"log"
	"todo-app"
	"todo-app/package/handler"
)

func main() {
	srv := new(todo.Server)
	handlers := new(handler.Handler)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("an error occurred while running the server:\n\t%s", err.Error())
	}

} 