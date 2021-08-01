package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/config"
	"rahmanfaisal10/embrio4-service/pkg/util"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

func DBConnection() *sqlx.DB {
	cfg := config.Get()

	connection := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=true&maxAllowedPacket=0",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBConnection,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sqlx.Connect("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(util.DBMAXCONNS)

	log.Info("connected to database")
	return db
}
