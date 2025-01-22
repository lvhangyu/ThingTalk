package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lvhangyu/ThingTalk/app/user/internal/data/query"
	"github.com/lvhangyu/ThingTalk/pkg/conf"
	"github.com/lvhangyu/ThingTalk/pkg/db/mysql"
	"github.com/redis/go-redis/v9"

	rdb "github.com/lvhangyu/ThingTalk/pkg/db/redis"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	db := mysql.Init(c.Database.Source)
	query.SetDefault(db)

	rdb := rdb.Init(c.Redis.Addr, "", 0)

	return &Data{
		db:  db,
		rdb: rdb,
	}, cleanup, nil
}
