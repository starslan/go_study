package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go_study/internal/app/config"
	"go_study/internal/app/file"
	"go_study/internal/app/httpserver/handlers"
	"log"
	"net/http"
)

func main() {

	cfg := config.NewConfig()
	shortURLList := make(map[string]string)
	if cfg.FilePath != "" {

		fillList(shortURLList, "shortURLList")
	}
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

	log.Fatal(http.ListenAndServe(cfg.ServerAddress, r))
}

func fillList(shortURLList map[string]string, path string) {

	cons, err := file.NewConsumer(path)
	if err != nil {
		panic(err)
	}

	for {
		event, err := cons.ReadEvent()
		if err != nil || event == nil {
			break
		}
		shortURLList[event.ID] = event.URL
	}
}
