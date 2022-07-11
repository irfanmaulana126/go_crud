package config

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

type ContextKey string
type Config struct {
	AppEnv                  string `mapstructure:"APPENV"`
	DBDriver                string `mapstructure:"DB_DRIVER"`
	DBSource                string `mapstructure:"DB_SOURCE"`
	DBMaxOpenConnections    int    `mapstructure:"DB_MAX_OPEN_CONNECTIONS"`
	DBMaxIdleConnections    int    `mapstructure:"DB_MAX_IDLE_CONNECTIONS"`
	AppIsDev                bool
	ServerAddress           string `mapstructure:"SERVER_ADDRESS"`
	JwtSigningKey           string `mapstructure:"JWT_SINGNING_KEY"`
	JwtIssuer               string `mapstructure:"JWT_ISSUER"`
	JwtAccessTokenDuration  int    `mapstructure:"JWT_ACCESS_TOKEN_DURATION_SECONDS"`
	JwtRefreshTokenDuration int    `mapstructure:"JWT_REFRESH_TOKEN_DURATION_SECONDS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func NewConfig() (*Config, error) {
	env := os.Getenv("APPENV")
	if env == "" {
		env = "app"
	}

	viper.AddConfigPath("./package/config")
	viper.SetConfigName(env)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SetConfigName("placeholder")

			if err := viper.ReadInConfig(); err != nil {
				log.Error().Err(err).Msg("[NewConfig-1] Failed To Read Config")
				return nil, err
			}
		} else {
			log.Error().Err(err).Msg("[NewConfig-2] Failed To Read Config")
			return nil, err
		}
	}

	cfg := &Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Error().Err(err).Msg("[NewConfig-3] Failed To Unmarshal Config")
		return nil, err
	}

	cfg.AppIsDev = cfg.AppEnv == "staging" || cfg.AppEnv == "local" || cfg.AppEnv == "dev"

	return cfg, nil
}
