package main

import (
	"log"
	"os"
	"to-do-list/database"
	"to-do-list/routes"

	"github.com/joho/godotenv"
)

type Todos struct {
	ID   int
	Task string
}

type TodoList struct {
	Todo []Todos
}

func main() {
	database.InitDb()

	router := routes.SetupRoutes()

	err := godotenv.Load()

	if err != nil {
		log.Printf("Error while Loading env")
	}

	port := os.Getenv("PORT")
	if err := router.Run(":" + port); err != nil {
		log.Printf("Error while Running Router")
	}

	log.Print("Server Running on port " + port)
}
