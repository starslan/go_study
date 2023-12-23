package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go_study/internal/app/httpserver/handlers"
	"net/http"
)

func main() {

	shortURLList := make(map[string]string)
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/", func(r chi.Router) {
		r.Post("/", handlers.ShortURLHandler(shortURLList))
		r.Get("/{id}", handlers.ShortURLHandler(shortURLList))
		r.Post("/api/shorten", handlers.ShortenURLHandler(shortURLList))
	})
	//http.HandleFunc(, handlers.ShortURLHandler(shortURLList))
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
