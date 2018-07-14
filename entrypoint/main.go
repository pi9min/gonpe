package entrypoint

import (
	"net/http"

	"github.com/ponpe/server/handler"
)

func init() {
	http.Handle("/", handler.Init())
}
