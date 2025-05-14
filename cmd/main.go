package main

import (
	"log"
	todoapp "todolist"
	"todolist/config"
	"todolist/internal/handler"
	"todolist/internal/repository"
	"todolist/internal/service"
)

func main() {
	config := config.InitConfig()
	log.Println("Config: ", config)

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoapp.Server)
	err := srv.Run(config.Port, handlers.InitRoutes())
	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}
