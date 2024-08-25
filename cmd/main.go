package main

import (
	"bug_busters/api"
	"bug_busters/internal/service"
	"bug_busters/internal/storage/postgres"
	"bug_busters/pkg/config"
	logger2 "bug_busters/pkg/logger"
	"log"
)

func main() {
	cfg := config.Load()
	logger := logger2.NewLogger()

	db, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		logger.Error("error in connection", "error", err)
		log.Fatal(err)
	}

	auth := service.NewAuthService(postgres.NewAuthRepo(db), logger)
	ii := service.NewIIService(postgres.NewIIRepo(db), logger)
	user := service.NewUserService(logger, postgres.NewUserRepo(db))
	servs := service.NewService(postgres.NewServiceRepo(db))
	router := api.NewRouter(auth, ii, user, servs)
	err = router.Run(cfg.GIN_PORT)

	if err != nil {
		logger.Error("error in router", "error", err)
		log.Fatal(err)
	}
}
