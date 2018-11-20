package handler

import (
	"context"
	"net/http"

	"github.com/ponpe/server/proto"

	"github.com/go-chi/chi"
	"github.com/ponpe/server/application"
)

type Handler struct {
	authApp *application.AuthenticationApp
}

func NewHandler(authApp *application.AuthenticationApp) http.Handler {
	h := &Handler{
		authApp: authApp,
	}
	hcHandler := pb.NewHealthCheckHandler(h)

	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		// Health check
		r.Route("/HealthCheck", func(r chi.Router) {
			r.HandleFunc("/Ping", hcHandler.Ping(nil))
		})
	})

	return r
}

func (h *Handler) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingResp, error) {
	return &pb.PingResp{
		Text: "Ping",
	}, nil
}
