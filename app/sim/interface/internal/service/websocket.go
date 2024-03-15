package service

import (
	"context"
	v1 "kratos_sim/api/sim/interface/v1"
)

func (s *SimService) Chat(ctx context.Context, in *v1.ChatReq) (*v1.ChatReply, error) {
	return s.ic.Chat(ctx, in)
}
