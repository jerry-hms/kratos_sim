package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos_sim/api/sim/interface/v1"
	"kratos_sim/app/sim/interface/internal/biz"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewSimService)

type SimService struct {
	v1.UnimplementedSimServer
	uc  *biz.UserUseCase
	ac  *biz.AuthUseCase
	ic  *biz.IMUseCase
	log log.Logger
}

func NewSimService(uc *biz.UserUseCase, ac *biz.AuthUseCase, ic *biz.IMUseCase, logger log.Logger) *SimService {
	return &SimService{
		uc:  uc,
		ac:  ac,
		ic:  ic,
		log: logger,
	}
}
