package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golobe/internal/controllers"
	"golobe/internal/model"
)

func HotelRoute(DB *sql.DB, router *gin.Engine) {
	hotelMethods := controllers.HotelScheme{
		DB:    DB,
		Hotel: model.Hotel{},
	}

	router.GET("/hotel", hotelMethods.GetHotels)
	router.GET("/hotel/:id", hotelMethods.GetHotelById)
	router.PATCH("/hotel/:id", hotelMethods.UpdateHotel)
	router.POST("/hotel", hotelMethods.CreateHotel)
	router.DELETE("/hotel/:id", hotelMethods.DeleteHotel)
}
