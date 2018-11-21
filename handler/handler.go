package handler

import (
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
	hcConv := pb.NewHealthCheckHTTPConverter(h)
	adminConv := pb.NewAdminHTTPConverter(h)

	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		// HealthCheck
		r.Route("/HealthCheck", func(r chi.Router) {
			r.Post("/Ping", hcConv.Ping(withErrorLog))
		})

		// Admin
		r.Route("/Admin", func(r chi.Router) {
			r.Post("/GetAllUser", adminConv.GetAllUser(withErrorLog))
			r.Post("/CreateGuestUser", adminConv.CreateGuestUser(withErrorLog))
		})
	})

	return r
}
