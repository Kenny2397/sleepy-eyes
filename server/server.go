package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Config struct {
	PORT         string
	JWT_SECRET   string
	DATABASE_URL string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *chi.Mux
}

// receiver function Broker es uparte del club Server
func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.PORT == "" {
		return nil, errors.New("port is required")
	}

	if config.DATABASE_URL == "" {
		return nil, errors.New("database is required")
	}

	if config.JWT_SECRET == "" {
		return nil, errors.New("jwt is required")
	}

	broker := &Broker{
		config: config,
		router: chi.NewRouter(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *chi.Mux)) {
	b.router = chi.NewRouter()

	binder(b, b.router)

	log.Println("Starting server on Port", b.Config().PORT)
	if err := http.ListenAndServe(b.config.PORT, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
