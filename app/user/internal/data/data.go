package data

import (
	"Kratos-demo/app/user/internal/data/query"
	"Kratos-demo/pkg/conf"
	"Kratos-demo/pkg/db/mysql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
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
