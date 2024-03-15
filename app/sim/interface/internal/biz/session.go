package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos_sim/api/session/service/v1"
)

type Session struct {
	Id                  int64
	UserId              int64
	RelId               int64
	SessionName         string
	SepId               int64
	LastSenderInfo      *v1.LastSenderInfo
	LastMessage         string
	UnreadMessageNumber int32
}

type Relation struct {
	Id         int64
	UserId     int64
	RelationId int64
	Scene      string
	SepId      int64
	CreatedAt  string
	UpdatedAt  string
}

type SessionRepo interface {
	CreateRelation(context.Context, *Relation) (*v1.CreateRelationReply, error)
	CreateSession(context.Context, *Session) (*v1.CreateSessionReply, error)
	ToFriend(ctx context.Context, me *Session, friend *Session) (meSes *v1.CreateSessionReply, friendSes *v1.CreateSessionReply, err error)
}

type SessionUseCase struct {
	repo SessionRepo
	log  *log.Helper
}

func NewSessionUseCase(repo SessionRepo, logger log.Logger) *SessionUseCase {
	return &SessionUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/session")),
	}
}
