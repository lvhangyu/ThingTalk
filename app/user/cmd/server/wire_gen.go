// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lvhangyu/ThingTalk/app/user/internal/biz"
	"github.com/lvhangyu/ThingTalk/app/user/internal/data"
	"github.com/lvhangyu/ThingTalk/app/user/internal/server"
	"github.com/lvhangyu/ThingTalk/app/user/internal/service"
	"github.com/lvhangyu/ThingTalk/pkg/conf"
	"github.com/lvhangyu/ThingTalk/pkg/grpc"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, bootstrap *conf.Bootstrap) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	cli := grpc.NewGrpcCli(bootstrap, logger)
	userRepo := data.NewUserRepo(dataData, logger, cli)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	userService := service.NewUserService(userUsecase)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
