package biz

import (
	"context"

	"Kratos-demo/pkg/grpc"
	"github.com/go-kratos/kratos/v2/log"
)

// User is a User model.
type User struct {
	Hello string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	ListByHello(context.Context, string) ([]*User, error)
	ListAll(context.Context) ([]*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo    UserRepo
	log     *log.Helper
	grpcCli *grpc.Cli
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, grpcCli *grpc.Cli, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger), grpcCli: grpcCli}
}

// CreateUser creates a User, and returns the new User.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", g.Hello)
	test, err := uc.grpcCli.UserGrpcClient.Test(ctx, nil)
	if err != nil {
		return nil, err
	}

	g.Hello = test.Message
	return uc.repo.Save(ctx, g)
}
