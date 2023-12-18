package main

import (
	"github.com/go-chi/chi/v5"
	"go_study/internal/app/httpserver/handlers"
	"net/http"
)

func main() {

	shortURLList := make(map[string]string)
	r := chi.NewRouter()
	r.Post("/", handlers.ShortURLHandler(shortURLList))
	r.Get("/", handlers.ShortURLHandler(shortURLList))
	//http.HandleFunc(, handlers.ShortURLHandler(shortURLList))
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
