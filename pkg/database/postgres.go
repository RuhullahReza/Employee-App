package database

import (
	"log"
	"os"
	"time"

	config "github.com/RuhullahReza/Employee-App/config"
	zlogr "github.com/RuhullahReza/Employee-App/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(cfg *config.Config) *gorm.DB {
	dsn := cfg.DSN()

	gormConfig := &gorm.Config{}
	gormConfig.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		zlogr.Log.Error(err, "failed to connect to database")
		panic("[ERROR] failed to connect to database")
	}

	return db
}
