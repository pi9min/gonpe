package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Init() http.Handler {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.HandleFunc("/Ping", func(w http.ResponseWriter, _ *http.Request) {
			w.Write([]byte("Ping"))
		})
	})

	return r
}
