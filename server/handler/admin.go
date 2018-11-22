package handler

import (
	"context"
	"time"

	"github.com/pi9min/gonpe/server/proto"
)

func (h *Handler) GetAllUser(ctx context.Context, req *pb.GetAllUserReq) (*pb.GetAllUserResp, error) {
	users, err := h.adminApp.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}

	pUsers := make([]*pb.User, 0, len(users))
	for i := range users {
		pu, err := users[i].Proto()
		if err != nil {
			return nil, err
		}

		pUsers = append(pUsers, pu)
	}

	return &pb.GetAllUserResp{
		Users: pUsers,
	}, nil
}

func (h *Handler) CreateGuestUser(ctx context.Context, req *pb.CreateGuestUserReq) (*pb.CreateGuestUserResp, error) {
	now := time.Now()
	if err := h.adminApp.CreateUser(ctx, req.Email, req.Password, now); err != nil {
		return nil, err
	}

	return &pb.CreateGuestUserResp{}, nil
}
