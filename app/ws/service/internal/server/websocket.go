package server

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_sim/api/ws/service/v1"
	"kratos_sim/app/ws/service/internal/conf"
	"kratos_sim/app/ws/service/internal/pkg/websocket"
	"kratos_sim/app/ws/service/internal/service"
)

func NewWebsocketServer(c *conf.Server, service *service.WsService, logger log.Logger) *websocket.Server {
	srv := websocket.NewServer(
		websocket.WithAddress(c.Websocket.Addr),
		websocket.WithPath(c.Websocket.Path),
		websocket.WithConnectHandle(service.OnWebsocketConnect),
		websocket.WithCodec("json"),
		websocket.WithPayloadType(websocket.PayloadTypeText),
	)

	service.SetWebsocketServer(srv)

	srv.RegisterMessageHandler(websocket.MessageType(v1.MessageType_Chat),
		func(sessionID websocket.SessionID, payload websocket.MessagePayload) error {
			switch t := payload.(type) {
			case *v1.SendReq:
				return service.OnMessage(sessionID, t)
			default:
				return errors.New("invalid payload type")
			}
		},
		func() websocket.Any {
			return &v1.SendReq{}
		},
	)
	return srv
}
