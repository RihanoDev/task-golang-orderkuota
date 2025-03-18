package config

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

var (
	db   *sql.DB
	once sync.Once
)

func InitDB() (*sql.DB, error) {
	var err error
	once.Do(func() {
		LoadENV()
		dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
			GetConfig("DB_USERNAME"),
			GetConfig("DB_PASSWORD"),
			GetConfig("DB_HOST"),
			GetConfig("DB_PORT"),
		)

		db, err = sql.Open("mysql", dsnWithoutDB)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to database server")
			return
		}

		dbName := GetConfig("DB_NAME")
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create database")
			return
		}

		db.Close()

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			GetConfig("DB_USERNAME"),
			GetConfig("DB_PASSWORD"),
			GetConfig("DB_HOST"),
			GetConfig("DB_PORT"),
			dbName,
		)

		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to database")
			return
		}

		if err = db.Ping(); err != nil {
			log.Fatal().Err(err).Msg("Database is unreachable")
			return
		}

		log.Info().Msg("Database connected successfully")
	})

	return db, err
}

func GetDB() *sql.DB {
	return db
}
