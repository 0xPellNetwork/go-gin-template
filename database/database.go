package database

import (
	"gin-template/config"
	"gin-template/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(cfg config.DatabaseConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch cfg.Driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.DSN), &gorm.Config{})
	default:
		db, err = gorm.Open(sqlite.Open(cfg.DSN), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	// 自动迁移模型
	err = db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
