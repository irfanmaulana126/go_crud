package manager

import (
	errorHelper "belajar/helper/error"
	paginationHelper "belajar/helper/pagination"
	jwtAuth "belajar/package/auth/jwt"
	middlewareAuth "belajar/package/auth/middleware"
	"belajar/package/config"
	gormDatabase "belajar/package/database/gorm"
	mysqlDatabase "belajar/package/database/mysql"
	"belajar/package/server"
	"database/sql"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Manager interface {
	GetConfig() *config.Config
	GetServer() *server.Server
	GetMysql() *sql.DB
	GetGorm() *gorm.DB
	GetJwt() jwtAuth.Jwt
	GetMiddleware() middlewareAuth.Middleware
	GetPagination() paginationHelper.Pagination
	GetMessageError() errorHelper.ErrorMessage
}

type manager struct {
	config         *config.Config
	server         *server.Server
	dbMysql        *sql.DB
	dbGorm         *gorm.DB
	jwtAuth        jwtAuth.Jwt
	middlewareAuth middlewareAuth.Middleware
	pagination     paginationHelper.Pagination
	messageError   errorHelper.ErrorMessage
}

func NewInit() (Manager, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Error().Err(err).Msg("[NewInit-1] Failed to Initialize Configuration")
		return nil, err
	}

	srv := server.NewServer(cfg)

	dbMysql, err := mysqlDatabase.NewMysql(cfg).Connect()
	if err != nil {
		log.Error().Err(err).Msg("[NewInit-2] Failed to Initialize Database Mysql")
		return nil, err
	}

	dbGorm, err := gormDatabase.NewGorm(cfg).Connect()
	if err != nil {
		log.Error().Err(err).Msg("[NewInit-3] Failed to Initialize Database Gorm")
		return nil, err
	}

	log.Logger = log.With().Caller().Logger()
	if cfg.AppIsDev {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}).With().Caller().Logger()
	}

	paginationHelper := paginationHelper.NewPagination()
	jwt := jwtAuth.NewJwt(cfg)
	middleware := middlewareAuth.NewMiddleware(cfg)

	errMessage := errorHelper.NewErrorMessage()
	err = errMessage.InitJSON("error.json")
	if err != nil {
		log.Error().Err(err).Msg("[NewInit-6] Failed to Initialize config error JSON file")
		return nil, err
	}

	return &manager{
		config:         cfg,
		server:         srv,
		dbMysql:        dbMysql,
		dbGorm:         dbGorm,
		jwtAuth:        jwt,
		middlewareAuth: middleware,
		pagination:     paginationHelper,
		messageError:   errMessage,
	}, nil
}

func (sm *manager) GetConfig() *config.Config {
	return sm.config
}

func (sm *manager) GetServer() *server.Server {
	return sm.server
}

func (sm *manager) GetMysql() *sql.DB {
	return sm.dbMysql
}

func (sm *manager) GetGorm() *gorm.DB {
	return sm.dbGorm
}

func (sm *manager) GetJwt() jwtAuth.Jwt {
	return sm.jwtAuth
}

func (sm *manager) GetMiddleware() middlewareAuth.Middleware {
	return sm.middlewareAuth
}

func (sm *manager) GetPagination() paginationHelper.Pagination {
	return sm.pagination
}

func (sm *manager) GetMessageError() errorHelper.ErrorMessage {
	return sm.messageError
}
