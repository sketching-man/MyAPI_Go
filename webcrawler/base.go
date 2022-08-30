package webcrawler

import (
	"database/sql"
	"fmt"
)

const (
	DB_HOST = "localhost"
	DB_NAME = "crawler"
	DB_USER = "postgres"
	DB_PW   = "postgres"
)

var (
	db *sql.DB
)

func init() {
	dbInfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_USER, DB_PW, DB_NAME)

	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
}

func Start() {
	StartCrawling()
}

func Close() {
	if db != nil {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}
}
