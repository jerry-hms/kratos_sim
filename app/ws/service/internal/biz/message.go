package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos_sim/api/ws/service/v1"
)

type Message struct {
	Id          int64
	RelId       int64
	MessageId   string
	SepId       int64
	SenderId    int64
	Sender      []byte
	ReceiverId  int64
	SendContent []byte
	IsRead      int32
}

type MessageRepo interface {
	Create(ctx context.Context, message *Message) (*Message, error)
}

type MessageUseCase struct {
	repo MessageRepo
	log  *log.Helper
}

func NewMessageUseCase(repo MessageRepo, logger log.Logger) *MessageUseCase {
	return &MessageUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/message")),
	}
}

func (m *MessageUseCase) CreateMessage(ctx context.Context, req *v1.CreateMessageReq) (*Message, error) {
	return m.repo.Create(ctx, &Message{
		RelId:       req.RelId,
		MessageId:   req.MessageId,
		SepId:       req.SepId,
		SenderId:    req.SenderId,
		Sender:      []byte(req.Sender),
		ReceiverId:  req.ReceiverId,
		SendContent: []byte(req.SendContent),
		IsRead:      req.IsRead,
	})
}
