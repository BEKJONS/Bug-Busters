package main

import (
	"bug_busters/api"
	"bug_busters/internal/service"
	"bug_busters/internal/storage/postgres"
	"bug_busters/pkg/config"
	logger2 "bug_busters/pkg/logger"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
)

func main() {
	cfg := config.Load()

	logger := logger2.NewLogger()
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	CasbinEnforcer, err := casbin.NewEnforcer(path+"/pkg/casbin/model.conf", path+"/pkg/casbin/policy.csv")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	db, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		logger.Error("error in connection", "error", err)
		log.Fatal(err)
	}

	auth := service.NewAuthService(postgres.NewAuthRepo(db), logger)
	ii := service.NewIIService(postgres.NewIIRepo(db), logger)
	user := service.NewUserService(logger, postgres.NewUserRepo(db))
	servs := service.NewService(postgres.NewServiceRepo(db))
	sw := service.NewSWStorage(postgres.NewSWRepo(db), logger)
	router := api.NewRouter(auth, ii, user, servs, sw, CasbinEnforcer)
	err = router.Run(cfg.GIN_PORT)

	if err != nil {
		logger.Error("error in router", "error", err)
		log.Fatal(err)
	}
}
