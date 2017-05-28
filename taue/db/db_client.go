package db

import (
	"log"
	"fmt"
	"database/sql"
)

func OpenDB() (db *sql.DB, err error) {

	const (
		DB_USER     = ""
		DB_PASSWORD = ""
		DB_NAME     = ""
		DB_HOST     = ""
		DB_PORT     = ""
	)

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT)

	db, err = sql.Open("postgres", dbInfo)

	if err != nil {
		log.Fatalf("Error opening database: %q", err)
		return nil, err
	}

	return db, err

}
