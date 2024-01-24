package middleware

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func GzipMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
		}

		fmt.Println("gzip middleware!")

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)

	}
}

func GetPayloadRequest(w http.ResponseWriter, r *http.Request) []byte {
	var reader io.ReadCloser
	var err error

	if r.Header.Get("Content-Encoding") == "gzip" {

		reader, err = gzip.NewReader(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			defer reader.Close()
		}
	} else {
		reader = r.Body
	}
	defer r.Body.Close()

	payload, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}

	return payload
}
