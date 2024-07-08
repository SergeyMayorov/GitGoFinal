package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SergeyMayorov/GitGoFinal/pkg/handlers"
	"github.com/SergeyMayorov/GitGoFinal/pkg/repository"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("APP_PORT")
	port = fmt.Sprintf(":%s", port)

	db := repository.New(repository.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	})
	app := handlers.New(db, port)

	err = app.Start(handlers.Routes(app))
	if err != nil {
		log.Fatal("start application failed %v", err)
	}
}
