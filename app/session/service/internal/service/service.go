package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "kratos_sim/api/session/service/v1"
	"kratos_sim/app/session/service/internal/biz"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewSessionService)

type SessionService struct {
	v1.UnimplementedSessionServer
	sc  *biz.SessionUseCase
	rc  *biz.RelationUseCase
	log *log.Helper
}

func NewSessionService(sc *biz.SessionUseCase, rc *biz.RelationUseCase, logger log.Logger) *SessionService {
	return &SessionService{
		sc:  sc,
		rc:  rc,
		log: log.NewHelper(log.With(logger, "module", "server/session")),
	}
}
