package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"stream-radar/domain/model"
	"stream-radar/internal/config"
	"stream-radar/internal/logger"
)

var DB *gorm.DB

func Connect(dbConfig config.DbConfig) {
	log := logger.GetInstance()
	dsn := fmt.Sprintf(
		"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Db,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	log.Debug("Connection Opened to Database")

	// Migrate the schemas
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("error migrating database")
	}
	log.Debug("Database Migrated")
	DB = db

}
