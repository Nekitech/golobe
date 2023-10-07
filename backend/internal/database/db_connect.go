package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func ConnectDB() *sql.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s",
		host, user, dbname, password, port, sslmode)

	db, err := sql.Open("postgres", dbURI)

	CheckError(err)

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Success connect to db !!!")
	}

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
