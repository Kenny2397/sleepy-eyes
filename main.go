package main

import (
	"context"
	"log"
	"net/http"

	"os"

	"github.com/Kenny2397/visual-programming/handlers"
	"github.com/Kenny2397/visual-programming/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link", "Access-Control-Allow-Headers", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// handler := cors.Default().Handler(r)

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		// cors.Default().Handler,
	)

	// r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.Get("/", handlers.HomeHandler(s))
	r.Post("/new", handlers.CreateDrawflow(s))
	r.Get("/drawflows", handlers.ReadAll(s))
	r.Get("/drawflow/{if}", handlers.GetDrawflowByIdg(s))
	r.Delete("/drawflow/{if}", handlers.DeleteDrawflowByIdg(s))

	r.Post("/run", handlers.RunProgram(s))
	// r.Get("/execute/{pythonCode}", handlers.ExecuteProgram(s))
}
