package config

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/RuhullahReza/Employee-App/pkg/logger"

	"github.com/spf13/viper"
)

type Config struct {
	DbHost         string `mapstructure:"DB_HOST"         default:"localhost"`
	DbPort         string `mapstructure:"DB_PORT"         default:"5433"`
	DbName         string `mapstructure:"DB_NAME"         default:"employee-service-db"`
	DbUsername     string `mapstructure:"DB_USERNAME"     default:"root"`
	DbPassword     string `mapstructure:"DB_PASSWORD"     default:""`
	AppName        string `mapstructure:"APP_NAME"        default:"employee-service"`
	AppHost        string `mapstructure:"APP_HOST"        default:":8080"`
	EndpointPrefix string `mapstructure:"ENDPOINT_PREFIX" default:"/api"`
}

var config *Config
var ErrFailUnmarshal = errors.New("failed to unmarshal config")

func NewConfig() (*Config, error) {
	if config == nil {
		config = new(Config)

		viper.SetConfigType("yaml")
		viper.SetConfigFile("env.yaml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()

		_ = viper.ReadInConfig()

		e := reflect.ValueOf(config).Elem()
		t := e.Type()
		for i := 0; i < e.NumField(); i++ {
			key := t.Field(i).Tag.Get("mapstructure")
			value := t.Field(i).Tag.Get("default")
			viper.SetDefault(key, value)
		}

		err := viper.Unmarshal(config)
		if err != nil {
			logger.Log.Error(err, "failed to unmarshal config")
			return nil, ErrFailUnmarshal
		}
	}

	return config, nil
}

func (cfg Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUsername,
		cfg.DbPassword,
		cfg.DbName,
	)
}
