package main

import (
	"fmt"
	todoapp "todolist"
	"todolist/config"
	"todolist/internal/handler"
	"todolist/internal/repository"
	"todolist/internal/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	servConfig := config.InitConfig()
	fmt.Println("Config: ", servConfig)

	dbConfig := config.InitDbConfig()
	db, err := repository.NewPostgresDB(*dbConfig)
	if err != nil {
		logrus.Fatal("Failed to connect to database: ", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoapp.Server)
	err = srv.Run(servConfig.Port, handlers.InitRoutes())
	if err != nil {
		logrus.Fatal("Failed to start the server: ", err)
	}
}

// migrate -path ./schema -database "postgres://postgres:"qwerty"@localhost:5436/postgres?sslmode=disable" up
