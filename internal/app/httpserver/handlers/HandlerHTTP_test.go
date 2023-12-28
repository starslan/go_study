package handlers

import (
	"bytes"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"go_study/internal/app/config"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type args struct {
	shortURLList map[string]string
}

type want struct {
	statusCode   int
	responseBody string
	contentType  string
}

type req struct {
	method  string
	payload string
	path    string
}

type tests []struct {
	name string
	args args
	want want
	req  req
}

func cfg() config.Config {
	return config.NewConfig()
}

func TestShortURLHandler(t *testing.T) {
	//type args struct {
	//	shortURLList map[string]string
	//}
	//
	//type want struct {
	//	statusCode   int
	//	responseBody string
	//	contentType  string
	//}
	//
	//type req struct {
	//	method  string
	//	payload string
	//	path    string
	//}

	tests := []struct {
		name string
		args args
		want want
		req  req
	}{

		{
			name: "Test POST '/' #1.",
			args: args{shortURLList: map[string]string{}},
			want: want{
				statusCode:   http.StatusCreated,
				responseBody: cfg().BaseURL + "/1",
				contentType:  "text/plain",
			},
			req: req{
				method:  http.MethodPost,
				path:    "/",
				payload: "http://ya.ru",
			},
		},
		{
			name: "Test POST '/' #2.",
			args: args{shortURLList: map[string]string{"1": "https://google.com"}},
			want: want{
				statusCode:   http.StatusCreated,
				responseBody: cfg().BaseURL + "/2",
				contentType:  "text/plain",
			},
			req: req{
				method:  http.MethodPost,
				path:    "/",
				payload: "http://ya.ru",
			},
		},
		{
			name: "Test GET '/{id}' #3.",
			args: args{shortURLList: map[string]string{"1": "https://google.com"}},
			want: want{
				statusCode:   http.StatusTemporaryRedirect,
				responseBody: cfg().BaseURL + "/1",
				contentType:  "text/plain",
			},
			req: req{
				method:  http.MethodGet,
				path:    "/1",
				payload: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var request *http.Request
			w := httptest.NewRecorder()
			h := ShortURLHandler(tt.args.shortURLList, cfg())
			if tt.req.method == http.MethodPost {
				request = httptest.NewRequest(tt.req.method, tt.req.path, bytes.NewBufferString(tt.req.payload))
				h.ServeHTTP(w, request)

			} else {
				request = httptest.NewRequest(tt.req.method, tt.req.path+tt.req.payload, nil)
				r := chi.NewRouter()
				r.Get("/{id}", ShortURLHandler(tt.args.shortURLList, cfg()))
				r.ServeHTTP(w, request)

			}
			res := w.Result()

			if tt.req.method == http.MethodPost {
				var resultBuf bytes.Buffer
				defer res.Body.Close()
				if _, err := io.Copy(&resultBuf, res.Body); err != nil {
					panic(err)
				}

				assert.Equalf(t, resultBuf.String(), tt.want.responseBody, "The wait result  %s  not equal got %s !", tt.want.responseBody, resultBuf.String())
			}

			assert.Equalf(t, res.StatusCode, tt.want.statusCode, "The wait statusCode  %d  not equal got %d !", res.StatusCode, tt.want.statusCode)
		})
	}
}

func TestShortenURLHandler(t *testing.T) {

	tests := tests{
		{
			name: "Test shorten POST '/' #1.",
			args: args{shortURLList: map[string]string{}},
			want: want{
				statusCode:   http.StatusCreated,
				responseBody: "{\"result\":\"" + cfg().BaseURL + "/1\"}",
				contentType:  "pplication/json",
			},
			req: req{
				method:  http.MethodPost,
				path:    "/api/shorten",
				payload: "{\"url\":\"http://ya.ru\"}",
			},
		},
		{
			name: "Test shorten POST '/' #2.",
			args: args{shortURLList: map[string]string{"1": "https://google.com"}},
			want: want{
				statusCode:   http.StatusCreated,
				responseBody: "{\"result\":\"" + cfg().BaseURL + "/2\"}",
				contentType:  "pplication/json",
			},
			req: req{
				method:  http.MethodPost,
				path:    "/api/shorten",
				payload: "{\"url\":\"http://ya.ru\"}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var request *http.Request
			w := httptest.NewRecorder()
			h := ShortenURLHandler(tt.args.shortURLList, cfg())

			request = httptest.NewRequest(tt.req.method, tt.req.path, bytes.NewBufferString(tt.req.payload))
			h.ServeHTTP(w, request)
			res := w.Result()

			var resultBuf bytes.Buffer
			defer res.Body.Close()
			if _, err := io.Copy(&resultBuf, res.Body); err != nil {
				panic(err)
			}

			assert.Equalf(t, resultBuf.String(), tt.want.responseBody, "The wait result  %s  not equal got %s !", tt.want.responseBody, resultBuf.String())
		})
	}

}
