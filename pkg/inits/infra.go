package inits

import (
	"fmt"
	"mnc-rest-api/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitializeConfig() config.Config {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	return cfg
}

func InitializeDatabase(cfg config.DbConfig, timezone string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable TimeZone=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Schema,
		timezone,
	)

	db, err := gorm.Open(
		postgres.Open(
			dsn), &gorm.Config{
			// SkipDefaultTransaction: true,
			PrepareStmt: true,
			Logger:      logger.Default.LogMode(logger.Silent),
		})
	if err != nil {
		panic(err)
	}

	return db
}
