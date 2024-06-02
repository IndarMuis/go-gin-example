package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(config Config) *gorm.DB {
	// * USING SQL DRIVER *
	//db, err := sql.Open("mysql", "root:al0homora@tcp(localhost:3306)/gin_db")
	//if err != nil {
	//	fmt.Println("ERROR CONNECT TO DATABASE")
	//	panic(err)
	//}
	//db.SetMaxIdleConns(10)
	//db.SetMaxOpenConns(20)
	//db.SetConnMaxIdleTime(5 * time.Minute)
	//db.SetConnMaxLifetime(60 * time.Minute)
	//fmt.Println("SUCCESS CONNECT TO DATABASE")

	dbHost := config.Get("DB_HOST")
	dbPort := config.Get("DB_PORT")
	dbName := config.Get("DB_NAME")
	dbUser := config.Get("DB_USER")
	dbPassword := config.Get("DB_PASSWORD")
	// dbTimezone := config.Get("DB_TIMEZONE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("DB CONNECTION REFUSED")
		panic(err)
	}

	return db
}
