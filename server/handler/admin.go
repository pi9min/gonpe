package handler

import (
	"context"
	"time"

	"github.com/pi9min/gonpe/server/proto"
)

func (h *Handler) GetAllUser(ctx context.Context, req *pb.GetAllUserReq) (*pb.GetAllUserResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

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

func (h *Handler) ChangeRole(ctx context.Context, req *pb.ChangeRoleReq) (*pb.ChangeRoleResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	now := time.Now()
	if err := h.adminApp.ChangeRole(ctx, req.UserId, req.Role, now); err != nil {
		return nil, err
	}

	return &pb.ChangeRoleResp{}, nil
}
