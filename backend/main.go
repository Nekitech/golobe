package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golobe/controllers"
	"golobe/database"
	"golobe/model"
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

	//err := db.AutoMigrate(&model.Hotel{}, &model.Room{}, &model.User{}, &model.HistoryBooking{}, &model.Booking{})

	//if err != nil {
	//	panic(err)
	//}

	hotelControllers := controllers.HotelScheme{
		Model: gorm.Model{},
		DB:    db,
		Hotel: model.Hotel{},
	}

	router := gin.Default()
	router.GET("/hotel", hotelControllers.GetHotels)

	//routes.HotelRoute(db, router)
	//routes.RoomRoute(db, router)
	//routes.UserRoute(db, router)

	err := router.Run("localhost:8090")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server start on port::8090")

	// close database
	defer db.Close()

}
