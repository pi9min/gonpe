package handler

import (
	"context"

	"github.com/pi9min/gonpe/server/proto"
)

func (h *Handler) SignIn(ctx context.Context, req *pb.SignInReq) (*pb.SignInResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	token, err := h.authenticationApp.SignIn(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.SignInResp{
		AccessToken: token,
	}, nil
}
