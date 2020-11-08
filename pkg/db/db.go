package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// Config :
type Config struct {
	Driver  string
	DB      string
	SSLMode string
}

// GetDB : open the database
func GetDB(conf *Config) *sqlx.DB {
	dbx, err := sqlx.Connect(conf.Driver, conf.DB)
	if err != nil {
		log.Fatalln(err)
	}
	dbx.SetMaxOpenConns(50)

	return dbx
}
