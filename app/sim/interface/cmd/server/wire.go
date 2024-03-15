//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos_sim/app/sim/interface/internal/biz"
	"kratos_sim/app/sim/interface/internal/conf"
	"kratos_sim/app/sim/interface/internal/data"
	"kratos_sim/app/sim/interface/internal/server"
	"kratos_sim/app/sim/interface/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
