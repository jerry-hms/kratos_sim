package service

import (
	"context"
	v1 "kratos_sim/api/session/service/v1"
)

func (s *SessionService) CreateRelation(ctx context.Context, in *v1.CreateRelationReq) (*v1.CreateRelationReply, error) {
	rel, err := s.rc.CreateRelation(ctx, in)
	if err != nil {
		return nil, err
	}

	return &v1.CreateRelationReply{
		Id:         rel.Id,
		UserId:     rel.UserId,
		RelationId: rel.RelationId,
		Scene:      rel.Scene,
		SepId:      rel.SepId,
		CreatedAt:  rel.CreatedAt,
	}, nil
}
