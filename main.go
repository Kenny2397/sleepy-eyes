package main

import (
	"context"
	"log"
	"os"

	"github.com/Kenny2397/visual-programming/handlers"
	"github.com/Kenny2397/visual-programming/server"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// loading .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get environment variables
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	// CREATE SERVER
	s, err := server.NewServer(context.Background(), &server.Config{
		PORT:         PORT,
		JWT_SECRET:   JWT_SECRET,
		DATABASE_URL: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	// start server
	s.Start(BindRoutes)

}

func BindRoutes(s server.Server, r *chi.Mux) {
	// r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.Get("/", handlers.HomeHandler(s))
}
