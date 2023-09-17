package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golobe/database"
	"golobe/model"
	"golobe/routes"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db := database.ConnectDB()

	err := db.AutoMigrate(&model.Hotel{}, &model.Room{})

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	routes.HotelRoute(db, router)
	routes.RoomRoute(db, router)

	err = router.Run("localhost:8090")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server start on port::8090")

}
