package biz

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	sessionV1 "kratos_sim/api/session/service/v1"
	v1 "kratos_sim/api/sim/interface/v1"
	wsV1 "kratos_sim/api/ws/service/v1"
	myJwt "kratos_sim/app/sim/interface/internal/pkg/jwt"
	"sync"
)

var wg sync.WaitGroup

type SendMessage struct {
	SenderId   int64
	ReceiverId int64
	Scene      string
	MsgType    string
	Text       string
	Url        string
}

type Message struct {
	Id          int64
	RelId       int64
	MessageId   string
	SepId       int64
	SenderId    int64
	Sender      string
	ReceiverId  int64
	SendContent string
	IsRead      int32
}

type WsRepo interface {
	Send(context.Context, *User, *SendMessage) (*wsV1.SendReply, error)
	CreateMessage(ctx context.Context, message *Message) (*wsV1.CreateMessageReply, error)
}

type IMUseCase struct {
	wsRepo      WsRepo
	userRepo    UserRepo
	sessionRepo SessionRepo
	log         *log.Helper
}

func NewWsUseCase(repo WsRepo, userRepo UserRepo, sesRepo SessionRepo, logger log.Logger) *IMUseCase {
	return &IMUseCase{
		wsRepo:      repo,
		userRepo:    userRepo,
		sessionRepo: sesRepo,
		log:         log.NewHelper(log.With(logger, "module", "usecase/ws")),
	}
}

func (i *IMUseCase) Chat(ctx context.Context, req *v1.ChatReq) (*v1.ChatReply, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	wsSignal := make(chan *wsV1.SendReply)
	auth, _ := myJwt.GetAuth(ctx)
	// 获取发送人和接收人的信息
	user, err := i.userRepo.GetUser(ctx, &User{
		Id: auth.UserId,
	})
	if err != nil {
		return nil, err
	}
	friend, err := i.userRepo.GetUser(ctx, &User{
		Id: req.RecvId,
	})
	if err != nil {
		return nil, err
	}

	// 发送消息
	go func() {
		defer func() {
			wg.Done()
		}()
		var reply, err = i.wsRepo.Send(ctx,
			&User{
				Id:       user.Id,
				Nickname: user.Nickname,
				Avatar:   user.Avatar,
			},
			&SendMessage{
				SenderId:   user.Id,
				ReceiverId: req.RecvId,
				Scene:      req.Scene,
				MsgType:    req.MsgType,
				Text:       req.Text,
				Url:        req.Url,
			},
		)
		if err != nil {
			i.log.Error("消息发送失败:", err)
		}
		wsSignal <- reply
	}()

	go func() {
		defer func() {
			wg.Done()
		}()
		// 创建会话关系
		rel, err := i.sessionRepo.CreateRelation(ctx, &Relation{
			UserId:     user.Id,
			RelationId: friend.Id,
			Scene:      req.Scene,
			SepId:      0,
		})
		if err != nil {
			i.log.Error("会话关系创建失败:", err)
			return
		}
		select {
		case resp := <-wsSignal:
			if resp == nil {
				return
			}
			switch req.Scene {
			case "friend":
				// 创建会话
				senderInfo := &sessionV1.LastSenderInfo{
					UserId:   user.Id,
					Avatar:   user.Avatar,
					Nickname: user.Nickname,
				}
				_, _, err = i.sessionRepo.ToFriend(ctx,
					&Session{
						UserId:              auth.UserId,
						RelId:               rel.Id,
						SessionName:         friend.Nickname,
						SepId:               0,
						LastSenderInfo:      senderInfo,
						LastMessage:         resp.CtAs,
						UnreadMessageNumber: 0,
					},
					&Session{
						UserId:              req.RecvId,
						RelId:               rel.Id,
						SessionName:         user.Nickname,
						SepId:               0,
						LastSenderInfo:      senderInfo,
						LastMessage:         resp.CtAs,
						UnreadMessageNumber: 1,
					},
				)
				if err != nil {
					i.log.Error("会话创建失败:", err)
					return
				}
			case "group":
				// toGroup...
			}

			// 创建聊天记录
			userByte, _ := json.Marshal(&wsV1.SendReq_Sender{
				UserId:   user.Id,
				Nickname: user.Nickname,
				Avatar:   user.Avatar,
			})
			_, err = i.wsRepo.CreateMessage(ctx, &Message{
				RelId:       rel.Id,
				MessageId:   resp.MessageId,
				SenderId:    user.Id,
				Sender:      string(userByte),
				ReceiverId:  friend.Id,
				SendContent: resp.Content,
			})
			if err != nil {
				i.log.Error("消息创建失败：", err)
			}
		}
	}()
	wg.Wait()
	return &v1.ChatReply{}, nil
}
