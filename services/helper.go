package services

import (
	"database/sql"
	"fmt"
)

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3300)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
