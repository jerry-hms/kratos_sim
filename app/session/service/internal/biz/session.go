package biz

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos_sim/api/session/service/v1"
)

type Session struct {
	Id                  int64
	UserId              int64
	RelId               int64
	SessionName         string
	SepId               int64
	LastSenderInfo      []byte
	LastMessage         string
	UnreadMessageNumber int16
	CreatedAt           string
	UpdatedAt           string
}

type SessionRepo interface {
	GetSessionOrCreate(ctx context.Context, b *Session) (*Session, error)
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

func (s *SessionUseCase) CreateSession(ctx context.Context, req *v1.CreateSessionRequest) (*Session, error) {
	senderInfo, _ := json.Marshal(req.LastSenderInfo)
	newSession, err := s.repo.GetSessionOrCreate(ctx, &Session{
		UserId:              req.UserId,
		RelId:               req.RelId,
		SessionName:         req.SessionName,
		LastSenderInfo:      senderInfo,
		LastMessage:         req.LastMessage,
		UnreadMessageNumber: int16(req.UnreadMessageNumber),
	})
	if err != nil {
		return nil, err
	}

	return newSession, nil
}
