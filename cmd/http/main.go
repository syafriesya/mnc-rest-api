package http

import (
	"log"
	"mnc-rest-api/database"
	"mnc-rest-api/internal/domain"
	"mnc-rest-api/pkg/inits"
)

func Main() {

	cfg := inits.InitializeConfig()

	database.ConnectDB()

	db := inits.InitializeDatabase(cfg.Env.Db, cfg.Env.Tz)

	err := db.AutoMigrate(&domain.User{}, &domain.TopUp{}, &domain.Transaction{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}

	router := InitializeGin(cfg.Env.Server.Name, cfg.Env.Environment)

	InitializeRepositories(db)
	InitializeUsecases(cfg)
	InitializeControllers(router, cfg)

	startServer(router, cfg.Env.Server.Port)
}
