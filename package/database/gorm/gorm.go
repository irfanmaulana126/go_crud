package gormDatabase

import (
	"time"

	"belajar/package/config"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Gorm interface {
	Connect() (*gorm.DB, error)
}

type Options struct {
	master  string
	maxOpen int
	maxIdle int
}

func NewGorm(cfg *config.Config) Gorm {
	opt := new(Options)
	opt.master = cfg.DBSource
	opt.maxOpen = cfg.DBMaxOpenConnections
	opt.maxIdle = cfg.DBMaxIdleConnections

	return opt
}

func (o *Options) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(o.master), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("[Connect-1] Failed To Connect Gorm")
		return nil, err
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(o.maxOpen)
	sqlDB.SetMaxIdleConns(o.maxIdle)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
