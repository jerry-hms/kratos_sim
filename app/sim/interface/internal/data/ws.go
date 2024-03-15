package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_sim/app/sim/interface/internal/biz"
)
import wsV1 "kratos_sim/api/ws/service/v1"

type WsRepo struct {
	data *Data
	log  *log.Helper
}

func NewWsRepo(data *Data, logger log.Logger) *WsRepo {
	return &WsRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/ws")),
	}
}

func (w *WsRepo) Send(ctx context.Context, ub *biz.User, sb *biz.SendMessage) (*wsV1.SendReply, error) {
	return w.data.wc.Send(ctx, &wsV1.SendReq{
		ReceiverId: sb.ReceiverId,
		Scene:      sb.Scene,
		SenderId:   sb.SenderId,
		Sender: &wsV1.SendReq_Sender{
			UserId:   ub.Id,
			Nickname: ub.Nickname,
			Avatar:   ub.Avatar,
		},
		Msg: &wsV1.SendReq_Msg{
			Type: sb.MsgType,
			Text: sb.Text,
			Url:  sb.Url,
		},
	})
}

func (m *WsRepo) CreateMessage(ctx context.Context, message *biz.Message) (*wsV1.CreateMessageReply, error) {
	return m.data.wc.CreateMessage(ctx, &wsV1.CreateMessageReq{
		RelId:       message.RelId,
		MessageId:   message.MessageId,
		SepId:       message.SepId,
		SenderId:    message.SenderId,
		Sender:      message.Sender,
		ReceiverId:  message.ReceiverId,
		SendContent: message.SendContent,
		IsRead:      0,
	})
}
