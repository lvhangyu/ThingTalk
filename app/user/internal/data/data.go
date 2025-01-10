package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lvhangyu/ThingTalk/app/user/internal/data/query"
	"github.com/lvhangyu/ThingTalk/pkg/conf"
	"github.com/lvhangyu/ThingTalk/pkg/db/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	db := mysql.Init(c.Database.Source)
	query.SetDefault(db)

	return &Data{
		db: db,
	}, cleanup, nil
}
