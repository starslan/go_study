package main

import (
	"go_study/internal/app/httpServer"
	"net/http"
)

func main() {

	shortURLList := make(map[string]string)
	http.HandleFunc("/", httpServer.ShortURLHandler(shortURLList))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}
