package main

import (
	"api"
	"api/pkg/handler"
	"api/pkg/repository"
	"api/pkg/service"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("config %s, ", err)
	}
	db, err := repository.NewPostgres(repository.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		UserName: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	})
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	//handlers := new(handler.Handler)
	server := new(api.Server)
	err = server.Run("8080", handlers.InitRouter())
	if err != nil {
		log.Fatal(err)
	}

}
