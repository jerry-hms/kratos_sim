package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_sim/app/session/service/internal/biz"
)

type Session struct {
	NotDelModel
	UserId              int64 `json:"user_id" gorm:"user_id"`
	RelId               int64 `json:"rel_id" gorm:"rel_id"`
	SessionName         string
	SepId               int64
	LastSenderInfo      []byte
	LastMessage         string
	UnreadMessageNumber int16
}

func (s *Session) TableName() string {
	return "sim_im_session"
}

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

func (s *SessionRepo) GetSessionOrCreate(ctx context.Context, b *biz.Session) (*biz.Session, error) {
	session := Session{
		UserId:              b.UserId,
		RelId:               b.RelId,
		SessionName:         b.SessionName,
		LastSenderInfo:      b.LastSenderInfo,
		LastMessage:         b.LastMessage,
		UnreadMessageNumber: b.UnreadMessageNumber,
	}

	result := s.data.db.WithContext(ctx).
		Where("user_id = ?", b.UserId).
		Where("rel_id = ?", b.RelId).
		FirstOrCreate(&session)
	if result.Error != nil {
		return nil, result.Error
	}

	session.SessionName = b.SessionName
	session.LastSenderInfo = b.LastSenderInfo
	session.LastMessage = b.LastMessage
	session.UnreadMessageNumber = session.UnreadMessageNumber + b.UnreadMessageNumber

	result = s.data.db.WithContext(ctx).Save(&session)
	if result.Error != nil {
		return nil, result.Error
	}

	return &biz.Session{
		Id: int64(session.ID),
	}, nil
}
