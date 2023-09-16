package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	//sslmode := os.Getenv("sslmode")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s",
		host, user, dbname, password, port)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Success connect to db !!!")
	}

	return db
}
