package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos_sim/api/session/service/v1"
)

type Relation struct {
	Id         int64
	UserId     int64
	RelationId int64
	Scene      string
	SepId      int64
	CreatedAt  string
	UpdatedAt  string
}

type RelationRepo interface {
	GetRelationOrCreate(ctx context.Context, b *Relation) (*Relation, error)
}

type RelationUseCase struct {
	repo RelationRepo
	log  *log.Helper
}

func NewRelationUseCase(repo RelationRepo, logger log.Logger) *RelationUseCase {
	return &RelationUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/relation")),
	}
}

func (r *RelationUseCase) CreateRelation(ctx context.Context, req *v1.CreateRelationReq) (*Relation, error) {
	return r.repo.GetRelationOrCreate(ctx, &Relation{
		UserId:     req.UserId,
		RelationId: req.RelationId,
		Scene:      req.Scene,
		SepId:      0,
	})
}
