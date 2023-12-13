package main

import (
	"go_study/internal/app/http-server"
	"net/http"
)

func main() {

	shortURLList := make(map[string]string)
	http.HandleFunc("/", http_server.ShortURLHandler(shortURLList))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}
