package service

import (
	"context"
	"fmt"

	"Kratos-demo/app/user/internal/biz"
	v1 "Kratos-demo/pkg/pb/user/api/v1"
)

// UserService is a User service.
type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUsecase
}

// NewUserService new a User service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// SayHello implements helloworld.UserServer.
func (s *UserService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateUser(ctx, &biz.User{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

func (s *UserService) Test(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	fmt.Println("test ")
	return &v1.HelloReply{Message: "Hello " + "test..."}, nil
}
