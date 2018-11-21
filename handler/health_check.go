package handler

import (
	"context"

	"github.com/ponpe/server/proto"
)

func (h *Handler) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingResp, error) {
	return &pb.PingResp{
		Text: "Ping",
	}, nil
}
