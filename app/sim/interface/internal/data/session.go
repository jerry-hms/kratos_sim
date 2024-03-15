package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos_sim/api/session/service/v1"
	"kratos_sim/app/sim/interface/internal/biz"
)

type SessionRepo struct {
	data *Data
	log  *log.Helper
}

func NewSessionRepo(data *Data, logger log.Logger) *SessionRepo {
	return &SessionRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/session")),
	}
}

func (s *SessionRepo) CreateSession(ctx context.Context, b *biz.Session) (*v1.CreateSessionReply, error) {
	return s.data.sc.CreateSession(ctx, &v1.CreateSessionRequest{
		RelId:               b.RelId,
		SessionName:         b.SessionName,
		LastSenderInfo:      b.LastSenderInfo,
		LastMessage:         b.LastMessage,
		UnreadMessageNumber: b.UnreadMessageNumber,
		UserId:              b.UserId,
	})
}

func (s *SessionRepo) CreateRelation(ctx context.Context, rel *biz.Relation) (*v1.CreateRelationReply, error) {
	return s.data.sc.CreateRelation(ctx, &v1.CreateRelationReq{
		UserId:     rel.UserId,
		RelationId: rel.RelationId,
		Scene:      rel.Scene,
	})
}

func (s *SessionRepo) ToFriend(ctx context.Context, me *biz.Session, friend *biz.Session) (*v1.CreateSessionReply, *v1.CreateSessionReply, error) {
	meSes, err := s.CreateSession(ctx, me)
	if err != nil {
		return nil, nil, errors.New("创建己方会话失败")
	}

	friendSes, err := s.CreateSession(ctx, friend)
	if err != nil {
		return nil, nil, errors.New("创建对方会话失败")
	}

	return meSes, friendSes, nil
}

func (s *SessionRepo) ToGroup() {

}
