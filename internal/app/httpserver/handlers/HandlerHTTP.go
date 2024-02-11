package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go_study/internal/app/config"
	"go_study/internal/app/file"
	"go_study/internal/app/httpserver/middleware"
	"log"
	"net/http"
	"strconv"
)

var shortURLList map[string]URLUserItem
var cfg config.Config

type URLUserItem struct {
	URL    string `json:"url"`
	UserId uint32 `json:"user-id"`
}

type shortenBody struct {
	URL string `json:"url"`
}

type shortenResult struct {
	Result string `json:"result"`
}

func addShortURL(url []byte, shortURLList map[string]URLUserItem, userId uint32) string {

	var key = strconv.Itoa(len(shortURLList) + 1)
	//event := file.Event{ID: key, URL: string(url)}
	event := file.Event{ID: key, Data: file.URLUserItem(URLUserItem{string(url), userId})}
	if cfg.FilePath != "" {

		cons, err := file.NewProducer(cfg.FilePath)
		if err != nil {
			log.Println(err)
		}
		err = cons.WriteEvent(&event)
		if err != nil {
			log.Println(err)
		}
	}
	shortURLList[key] = URLUserItem{string(url), 1}
	return key
}

func getUserId(r *http.Request) uint32 {
	userId, err := strconv.ParseUint(r.Header.Get("X-USER-ID"), 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(userId)
}

func ShortURLHandler(shortURLList map[string]URLUserItem, config config.Config) http.HandlerFunc {
	cfg = config
	return func(w http.ResponseWriter, r *http.Request) {

		userId := getUserId(r)

		switch r.Method {
		case http.MethodGet:
			id := chi.URLParam(r, "id")
			if val, ok := shortURLList[id]; ok {
				w.Header().Set("Location", val.URL)
				w.WriteHeader(http.StatusTemporaryRedirect)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}

		case http.MethodPost:

			payload := middleware.GetPayloadRequest(w, r)

			w.WriteHeader(http.StatusCreated)
			var link = cfg.BaseURL + "/" + addShortURL(payload, shortURLList, userId)
			w.Write([]byte(link))

		}
	}

}

func ShortenURLHandler(shortURLList map[string]URLUserItem, config config.Config) http.HandlerFunc {
	cfg = config
	return func(w http.ResponseWriter, r *http.Request) {
		userId := getUserId(r)
		switch r.Method {
		case http.MethodPost:

			payload := middleware.GetPayloadRequest(w, r)

			value := shortenBody{}
			if err := json.Unmarshal(payload, &value); err != nil {
				panic(err)
			}

			var link = cfg.BaseURL + "/" + addShortURL([]byte(value.URL), shortURLList, userId)
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
