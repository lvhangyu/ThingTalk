package grpc

import (
	"context"

	"Kratos-demo/pkg/conf"
	v1 "Kratos-demo/pkg/pb/user/api/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGrpcCli)

type Cli struct {
	UserGrpcClient v1.UserClient
}

func NewGrpcCli(b *conf.Bootstrap, logger log.Logger) *Cli {
	return &Cli{UserGrpcClient: newUserClient(b, logger)}
}

func newUserClient(b *conf.Bootstrap, logger log.Logger) v1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(b.UserServer.Grpc.Addr),
		grpc.WithTimeout(b.UserServer.Grpc.Timeout.AsDuration()),
		grpc.WithMiddleware(
			circuitbreaker.Client(),
			recovery.Recovery(),
			validate.Validator(),
			//metadata.Client(o...),
			//logging.Client(logger),
		),
	)
	if err != nil {
		panic(err)
	}
	c := v1.NewUserClient(conn)
	return c
}
