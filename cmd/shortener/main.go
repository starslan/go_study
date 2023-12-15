package main

import (
	"go_study/internal/app/httpserver/handlers"
	"net/http"
)

func main() {

	shortURLList := make(map[string]string)
	http.HandleFunc("/", handlers.ShortURLHandler(shortURLList))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}
