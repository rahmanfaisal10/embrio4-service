package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/config"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func DBConnection() *sqlx.DB {
	cfg := config.Get()

	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBPassword, cfg.DBName)

	db, err := sqlx.Connect("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("connected to database")
	return db
}
