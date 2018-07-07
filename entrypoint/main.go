package entrypoint

import (
	"net/http"

	"github.com/go-chi/chi"
)

func init() {
	r := chi.NewRouter()

	// v1
	r.Route("/v1", func(r chi.Router) {
		r.HandleFunc("/HelloService.Ping", Ping)
	})

	http.Handle("/", r)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
