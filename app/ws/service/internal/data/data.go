package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"kratos_sim/app/ws/service/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMessageRepo)

type NotDelModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Data .
type Data struct {
	db *gorm.DB
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	logHelper := log.NewHelper(log.With(logger, "module", "ws-server/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Database.Prefix,
			SingularTable: true,
		},
	})
	if err != nil {
		logHelper.Fatalf("failed opening connection to mysql: %v", err)
	}

	return db
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: db,
	}, cleanup, nil
}
