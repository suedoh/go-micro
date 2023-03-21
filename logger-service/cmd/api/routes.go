package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
    mux := chi.NewRouter()

    mux.Use(cors.Handler(cors.Options{
        AllowedOrigins: []string{"http://*", "https://*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELTE", "OPTIONS"},
        AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
        ExposedHeaders: []string{"Link"},
        AllowCredentials: true,
        MaxAge: 300,
    }))

    mux.Use(middleware.Heartbeat("/ping"))

    mux.Post("/log", app.WriteLog)

    return mux
}
