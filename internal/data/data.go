package data

import (
	"hk591/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewHouseRepo, NewUserRepo)

// Data .
type Data struct {
	db *sqlx.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)

	db := sqlx.MustOpen(c.Database.Driver, c.Database.Source)
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)

	d := &Data{db: db}
	
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
