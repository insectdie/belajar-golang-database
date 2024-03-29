package belajargolangdatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
