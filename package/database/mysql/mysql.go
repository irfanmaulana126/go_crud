package mysqlDatabase

import (
	"database/sql"
	"time"

	"belajar/package/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

type Mysql interface {
	Connect() (*sql.DB, error)
}

type Options struct {
	master  string
	driver  string
	maxOpen int
	maxIdle int
}

func NewMysql(cfg *config.Config) Mysql {
	opt := new(Options)
	opt.driver = cfg.DBDriver
	opt.master = cfg.DBSource
	opt.maxOpen = cfg.DBMaxOpenConnections
	opt.maxIdle = cfg.DBMaxIdleConnections

	return opt
}

func (o *Options) Connect() (*sql.DB, error) {
	if o.driver == "" {
		o.driver = "mysql"
	}

	db, err := sql.Open(o.driver, o.master)
	if err != nil {
		log.Error().Err(err).Msg("[Connect-1] Failed To Connect Mysql")
		return nil, err
	}

	db.SetMaxOpenConns(o.maxOpen)
	db.SetMaxIdleConns(o.maxIdle)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
