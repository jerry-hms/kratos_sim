package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"kratos_sim/api/ws/service/v1"
	"kratos_sim/app/ws/service/internal/pkg/websocket"
	"kratos_sim/app/ws/service/internal/pkg/ws_msg"
)

func (s *WsService) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws
}

func (s *WsService) OnWebsocketConnect(session_id websocket.SessionID, register bool) {
	if register {
		msg := make(map[string]interface{})
		msg["code"] = 200
		msg["msg"] = "连接成功"
		msg["client_id"] = session_id
		err := s.ws.SendMessage(session_id, websocket.PayloadTypeString, msg)
		if err != nil {
			fmt.Println("消息发送失败", err)
			return
		}
	} else {
		// 下线...
	}
}

func (s *WsService) OnMessage(session_id websocket.SessionID, req interface{}) error {
	fmt.Printf("发送人:%s\n", session_id)
	fmt.Println("req", req)
	return nil
}

func (s *WsService) Bind(ctx context.Context, req *v1.BindReq) (*v1.BindReply, error) {
	s.client.Add(req.UserId, websocket.SessionID(req.ClientId))

	return &v1.BindReply{}, nil
}

func (s *WsService) Send(ctx context.Context, req *v1.SendReq) (*v1.SendReply, error) {
	session_id := s.client.Get(req.ReceiverId)

	msg := ws_msg.NewMessage().
		SetScene(req.Scene).
		SetSender(req.Sender).
		SetMessage(req.Msg).Parse()

	err := s.ws.SendMessage(session_id, websocket.PayloadTypeString, msg)
	if err != nil {
		return nil, errors.New("消息发送失败")
	}
	bytes, _ := json.Marshal(msg.GetMessage())
	ctAs, err := msg.GetCtAs()
	if err != nil {
		return nil, err
	}
	return &v1.SendReply{
		Content:   string(bytes),
		CtAs:      ctAs,
		MessageId: msg.GetMessageId(),
	}, nil
}
