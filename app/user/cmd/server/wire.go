//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Kratos-demo/app/user/internal/biz"
	"Kratos-demo/app/user/internal/data"
	"Kratos-demo/app/user/internal/server"
	"Kratos-demo/app/user/internal/service"
	"Kratos-demo/pkg/conf"
	"Kratos-demo/pkg/grpc"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, grpc.ProviderSet, newApp))
}
