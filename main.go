package main

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/config"
	"rahmanfaisal10/embrio4-service/pkg/handler"
	"rahmanfaisal10/embrio4-service/pkg/repository"
	"rahmanfaisal10/embrio4-service/pkg/service"

	"github.com/labstack/gommon/log"
)

func main() {
	//initialization db, repo, service and router connected
	cfg := config.Get()
	db := repository.DBConnection()
	repo := repository.InitRepository(db)
	svc := service.InitService(repo)
	router := handler.InitRouter(svc)

	//running system and start port
	err := router.Start(fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("server has started at port %s", cfg.Port)

	defer recoverServer()
}

func recoverServer() {
	if r := recover(); r != nil {
		log.Info("RECOVERED")
	}
}
