package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "kratos_sim/api/ws/service/v1"
	"kratos_sim/app/ws/service/internal/biz"
	"kratos_sim/app/ws/service/internal/pkg/websocket"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewWsService)

type WsService struct {
	v1.UnimplementedWsServer
	ws     *websocket.Server
	mc     *biz.MessageUseCase
	log    *log.Helper
	client *websocket.Client
}

func NewWsService(cli *websocket.Client, mc *biz.MessageUseCase, logger log.Logger) *WsService {
	return &WsService{
		mc:     mc,
		log:    log.NewHelper(log.With(logger, "module", "service/ws")),
		client: cli,
	}
}
