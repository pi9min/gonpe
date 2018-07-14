package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Init() http.Handler {
	r := chi.NewRouter()

	// v1
	r.Route("/v1", func(r chi.Router) {
		r.HandleFunc("/HelloService.Ping", Ping)
	})

	return r
}
