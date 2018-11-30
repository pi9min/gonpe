package handler

import (
	"context"
	"time"

	"github.com/pi9min/gonpe/server/proto"
)

func (h *Handler) RegisterGuestUser(ctx context.Context, req *pb.RegisterGuestUserReq) (*pb.RegisterGuestUserResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	now := time.Now()
	if err := h.authenticationApp.RegisterGuestUser(ctx, req.AuthProviderUserId, now); err != nil {
		return nil, err
	}

	return &pb.RegisterGuestUserResp{}, nil
}
