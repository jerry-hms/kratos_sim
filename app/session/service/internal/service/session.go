package service

import (
	"context"
	v1 "kratos_sim/api/session/service/v1"
)

func (s *SessionService) CreateSession(ctx context.Context, req *v1.CreateSessionRequest) (*v1.CreateSessionReply, error) {
	// 创建会话
	session, err := s.sc.CreateSession(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.CreateSessionReply{
		SessionId: session.Id,
	}, nil
}
