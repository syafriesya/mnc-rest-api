package database

import (
	"fmt"
	"mnc-rest-api/internal/config"
	"os"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db    *gorm.DB
	dbApm *gorm.DB
	err   error
)

// ConnectDB establishes a connection with PostGre server.
func ConnectDB() error {

	// models.ApiMap = make(map[string]models.Apis)
	if db, err = gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				config.Cfg.Env.Db.Host,
				config.Cfg.Env.Db.Port,
				config.Cfg.Env.Db.User,
				config.Cfg.Env.Db.Password,
				config.Cfg.Env.Db.Name)), &gorm.Config{
			PrepareStmt: true,
			Logger:      logger.Default.LogMode(logger.Silent),
		}); err != nil {
		return err
	}

	if err != nil {
		panic(err)
	}

	userType := os.Getenv("USER_TYPE")
	if userType == "local" {
		return nil
	} else {
		db.Logger = logger.Discard
		return nil
	}
}
