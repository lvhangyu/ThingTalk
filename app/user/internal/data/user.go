package data

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/lvhangyu/ThingTalk/app/user/internal/biz"
	"github.com/lvhangyu/ThingTalk/app/user/internal/data/model"
	"github.com/lvhangyu/ThingTalk/pkg/grpc"
	"gorm.io/gorm"
)

type userRepo struct {
	data    *Data
	log     *log.Helper
	grpcCli *grpc.Cli
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger, client *grpc.Cli) biz.UserRepo {
	return &userRepo{
		data:    data,
		log:     log.NewHelper(logger),
		grpcCli: client,
	}
}

func (r *userRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	err := r.data.db.Transaction(func(tx *gorm.DB) error {
		var code = "123"
		var price int64 = 1
		product := model.Product{
			ID:    0,
			Code:  &code,
			Price: &price,
		}
		tx.Create(&product)
		fmt.Println(product)

		_, err := r.grpcCli.UserGrpcClient.Test(ctx, nil)
		if err != nil {
			return err
		}

		return nil
		//return errors.New("test")
	})

	if err != nil {
		return nil, err
	}

	return g, nil

}

func (r *userRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	err := r.data.db.Transaction(func(tx *gorm.DB) error {
		p := model.Product{}
		err := tx.WithContext(ctx).Create(&p).Error
		if err != nil {
			return err
		}

		r.grpcCli.UserGrpcClient.Test(ctx, nil)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return g, nil
}

func (r *userRepo) FindByID(context.Context, int64) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) ListByHello(context.Context, string) ([]*biz.User, error) {
	return nil, nil
}

func (r *userRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}
