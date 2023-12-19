package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

var shortURLList map[string]string

func addShortURL(url []byte, shortURLList map[string]string) string {

	var key = strconv.Itoa(len(shortURLList) + 1)
	shortURLList[key] = string(url)
	return key
}

func ShortURLHandler(shortURLList map[string]string) http.HandlerFunc {
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
			var link = "http://" + r.Host + "/" + addShortURL(payload, shortURLList)
			w.Write([]byte(link))

		}
	}

}
