package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_sim/app/session/service/internal/biz"
)

type Relation struct {
	NotDelModel
	UserId     int64
	RelationId int64
	Scene      string
	SepId      int64
}

func (r *Relation) TableName() string {
	return "sim_im_session_relation"
}

type RelationRepo struct {
	data *Data
	log  *log.Helper
}

func NewRelationRepo(data *Data, logger log.Logger) *RelationRepo {
	return &RelationRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/relation")),
	}
}

func (r *RelationRepo) GetRelationOrCreate(ctx context.Context, b *biz.Relation) (*biz.Relation, error) {
	rel := Relation{
		UserId:     b.UserId,
		RelationId: b.RelationId,
		Scene:      b.Scene,
	}

	result := r.data.db.WithContext(ctx).FirstOrCreate(&rel)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Relation{
		Id:         int64(rel.ID),
		UserId:     rel.UserId,
		RelationId: rel.RelationId,
		Scene:      rel.Scene,
		SepId:      rel.SepId,
		CreatedAt:  rel.CreatedAt.Format("2006-01-02 15:04:05"),
	}, result.Error
}
