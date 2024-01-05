package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:al0homora@tcp(localhost:3306)/gin_db")
	if err != nil {
		fmt.Println("ERROR CONNECT TO DATABASE")
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	fmt.Println("SUCCESS CONNECT TO DATABASE")

	return db
}
