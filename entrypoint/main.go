package entrypoint

import (
	"net/http"
)

func init() {
	http.HandleFunc("/v1/HelloService.Ping", Ping)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
