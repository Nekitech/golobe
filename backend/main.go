package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golobe/database"
	"golobe/model"
	"golobe/structure"
	"gorm.io/gorm"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db := database.ConnectDB()

	db.AutoMigrate(&structure.Hotel{}, &structure.Room{})

	router := gin.Default()

	hotelMethods := model.HotelScheme{
		Model: gorm.Model{},
		DB:    db,
		Hotel: structure.Hotel{},
	}
	roomMethods := model.RoomScheme{
		Model: gorm.Model{},
		DB:    db,
		Room:  structure.Room{},
	}

	router.GET("/hotel", hotelMethods.GetHotels)
	router.GET("/hotel/:id", hotelMethods.GetHotelById)
	router.PATCH("/hotel/:id", hotelMethods.UpdateHotel)
	router.POST("/hotel", hotelMethods.CreateHotel)

	router.POST("/room", roomMethods.CreateRoom)

	err := router.Run("localhost:8090")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server start on port::8090")

}
