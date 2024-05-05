package database

import (
	"github.com/RuhullahReza/Employee-App/pkg/logger"

	"github.com/RuhullahReza/Employee-App/app/domain"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	logger.Log.Info("migrating database...")
	err := db.AutoMigrate(&domain.Employee{})
	if err != nil {
		logger.Log.Error(err, "database migration failed")
		panic("[ERROR] database migration failed")
	}
}