package db

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
	"github.com/labstack/gommon/log"
)

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./drinklytics.db") //creating the db connection
	if err != nil {
		log.Errorf("DB Connection Failed %v", err)
		return nil, err
	}
	return db, nil
}
