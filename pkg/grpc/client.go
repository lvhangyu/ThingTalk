package grpc

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/lvhangyu/ThingTalk/pkg/conf"
	device "github.com/lvhangyu/ThingTalk/pkg/pb/device/api/v1"
	user "github.com/lvhangyu/ThingTalk/pkg/pb/user/api/v1"
	"google.golang.org/grpc"
)

var ProviderSet = wire.NewSet(NewGrpcCli)

type Cli struct {
	UserGrpcClient   user.UserClient
	DeviceGrpcClient device.DeviceClient
}

func NewGrpcCli(b *conf.Bootstrap, logger log.Logger) *Cli {
	return &Cli{UserGrpcClient: newUserClient(b), DeviceGrpcClient: newDeviceClient(b)}
}

func newUserClient(b *conf.Bootstrap) user.UserClient {
	conn := newGrpcConn(b.UserServer.Grpc.Addr, b.UserServer.Grpc.Timeout.AsDuration())
	c := user.NewUserClient(conn)
	return c
}

func newDeviceClient(b *conf.Bootstrap) device.DeviceClient {
	conn := newGrpcConn(b.DeviceServer.Grpc.Addr, b.DeviceServer.Grpc.Timeout.AsDuration())
	c := device.NewDeviceClient(conn)
	return c
}

func newGrpcConn(addr string, timeout time.Duration) *grpc.ClientConn {
	conn, err := kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithEndpoint(addr),
		kgrpc.WithTimeout(timeout),
		kgrpc.WithMiddleware(
			circuitbreaker.Client(),
			recovery.Recovery(),
			validate.Validator(),
			metadata.Client(),
			//logging.Client(logger),
		),
	)
	if err != nil {
		panic(err)
	}

	return conn
}
