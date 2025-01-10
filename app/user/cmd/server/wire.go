//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lvhangyu/ThingTalk/app/user/internal/biz"
	"github.com/lvhangyu/ThingTalk/app/user/internal/data"
	"github.com/lvhangyu/ThingTalk/app/user/internal/server"
	"github.com/lvhangyu/ThingTalk/app/user/internal/service"
	"github.com/lvhangyu/ThingTalk/pkg/conf"
	"github.com/lvhangyu/ThingTalk/pkg/grpc"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, grpc.ProviderSet, newApp))
}
