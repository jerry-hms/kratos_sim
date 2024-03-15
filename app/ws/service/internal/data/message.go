package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_sim/app/ws/service/internal/biz"
)

type Message struct {
	NotDelModel
	RelId       int64
	MessageId   string
	SepId       int64
	SenderId    int64
	Sender      []byte
	ReceiverId  int64
	SendContent []byte
	IsRead      int32
}

func (m *Message) TableName() string {
	return "sim_im_session_message"
}

type MessageRepo struct {
	data *Data
	log  *log.Helper
}

func NewMessageRepo(data *Data, logger log.Logger) *MessageRepo {
	return &MessageRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/message")),
	}
}

func (m *MessageRepo) Create(ctx context.Context, b *biz.Message) (*biz.Message, error) {
	message := Message{
		RelId:       b.RelId,
		MessageId:   b.MessageId,
		SepId:       b.SepId,
		SenderId:    b.SenderId,
		Sender:      b.Sender,
		ReceiverId:  b.ReceiverId,
		SendContent: b.SendContent,
		IsRead:      0,
	}
	result := m.data.db.WithContext(ctx).Create(&message)
	return &biz.Message{
		Id: int64(message.ID),
	}, result.Error
}
