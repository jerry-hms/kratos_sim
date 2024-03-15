package service

import (
	"context"
	v1 "kratos_sim/api/sim/interface/v1"
)

func (s *SimService) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	return s.ac.Register(ctx, req)
}

func (s *SimService) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	return s.ac.Login(ctx, req)
}

func (s *SimService) BindWs(ctx context.Context, in *v1.BindWsReq) (*v1.BindWsReply, error) {
	return s.uc.Bind(ctx, in.ClientId)
}
