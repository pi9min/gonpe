package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pi9min/gonpe/server/application"
	"github.com/pi9min/gonpe/server/proto"
)

type Handler struct {
	adminApp          *application.AdminApp
	authenticationApp *application.AuthenticationApp
}

func NewHandler(
	adminApp *application.AdminApp,
	authenticationApp *application.AuthenticationApp,
) http.Handler {
	h := &Handler{
		adminApp:          adminApp,
		authenticationApp: authenticationApp,
	}
	healthCheckConv := pb.NewHealthCheckHTTPConverter(h)
	adminConv := pb.NewAdminHTTPConverter(h)
	authenticationConv := pb.NewAuthenticationHTTPConverter(h)

	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		// HealthCheck
		r.Post(csRestPath(healthCheckConv.PingWithName(withErrorLog)))

		// Admin
		r.Post(csRestPath(adminConv.GetAllUserWithName(withErrorLog)))
		r.Post(csRestPath(adminConv.ChangeRoleWithName(withErrorLog)))

		// Authentication
		r.Post(csRestPath(authenticationConv.RegisterGuestUserWithName(withErrorLog)))
	})

	return r
}

// csRestPath is make Case-Sensitive REST path.
// e.g. service: Foo, method: Bar
// -> "/Foo/Bar"
func csRestPath(service, method string, hfn http.HandlerFunc) (string, http.HandlerFunc) {
	return fmt.Sprintf("/%s/%s", service, method), hfn
}
