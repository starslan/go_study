package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go_study/internal/app/config"
	"go_study/internal/app/httpserver/handlers"
	"log"
	"net/http"
)

func main() {

	cfg := config.AppConfig()
	shortURLList := make(map[string]string)
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/", func(r chi.Router) {
		r.Post("/", handlers.ShortURLHandler(shortURLList, cfg))
		r.Get("/{id}", handlers.ShortURLHandler(shortURLList, cfg))
		r.Post("/api/shorten", handlers.ShortenURLHandler(shortURLList, cfg))
	})

	panic(cfg.BaseURL)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerAddress, r))
}
