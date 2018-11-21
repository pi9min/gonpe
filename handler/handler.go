package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ponpe/server/application"
	"github.com/ponpe/server/proto"
)

type Handler struct {
	adminApp *application.AdminApp
}

func NewHandler(adminApp *application.AdminApp) http.Handler {
	h := &Handler{
		adminApp: adminApp,
	}
	healthCheckConv := pb.NewHealthCheckHTTPConverter(h)
	adminConv := pb.NewAdminHTTPConverter(h)

	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		// HealthCheck
		r.Post(csRestPath(healthCheckConv.PingWithName(withErrorLog)))

		// Admin
		r.Post(csRestPath(adminConv.GetAllUserWithName(withErrorLog)))
		r.Post(csRestPath(adminConv.CreateGuestUserWithName(withErrorLog)))
	})

	return r
}

// csRestPath is make Case-Sensitive REST path.
// e.g. service: Foo, method: Bar
// -> "/Foo/Bar"
func csRestPath(service, method string, hfn http.HandlerFunc) (string, http.HandlerFunc) {
	return fmt.Sprintf("/%s/%s", service, method), hfn
}
