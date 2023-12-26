package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go_study/internal/app/config"
	"io"
	"log"
	"net/http"
	"strconv"
)

var shortURLList map[string]string

type shortenBody struct {
	URL string `json:"url"`
}

type shortenResult struct {
	Result string `json:"result"`
}

func addShortURL(url []byte, shortURLList map[string]string) string {

	var key = strconv.Itoa(len(shortURLList) + 1)
	shortURLList[key] = string(url)
	return key
}

//type MyHandler struct {
//	cfg config.Config
//}
//
//func NewMyHandler(cfg config.Config) *MyHandler {
//	return &MyHandler{cfg: cfg}
//}

func ShortURLHandler(shortURLList map[string]string, cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			id := chi.URLParam(r, "id")
			if val, ok := shortURLList[id]; ok {
				w.Header().Set("Location", val)
				w.WriteHeader(http.StatusTemporaryRedirect)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}

		case http.MethodPost:
			defer r.Body.Close()
			payload, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
			}

			w.WriteHeader(http.StatusCreated)
			fmt.Println(cfg)
			log.Println(cfg)
			var link = cfg.BaseURL + ":" + cfg.ServerAddress + "/" + addShortURL(payload, shortURLList)
			w.Write([]byte(link))

		}
	}

}

func ShortenURLHandler(shortURLList map[string]string, cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			defer r.Body.Close()

			payload, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
			}

			value := shortenBody{}
			if err := json.Unmarshal(payload, &value); err != nil {
				panic(err)
			}

			var link = cfg.BaseURL + ":" + cfg.ServerAddress + "/" + addShortURL([]byte(value.URL), shortURLList)
			sr := shortenResult{Result: link}

			//var resBody []byte
			resBody, err := json.Marshal(sr)
			if err != nil {
				panic(err)
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(resBody)

		}
	}

}
