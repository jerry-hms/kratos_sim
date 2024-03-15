package service

import (
	"context"
	v1 "kratos_sim/api/ws/service/v1"
)

func (w *WsService) CreateMessage(ctx context.Context, req *v1.CreateMessageReq) (*v1.CreateMessageReply, error) {
	_, err := w.mc.CreateMessage(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.CreateMessageReply{}, nil
}
