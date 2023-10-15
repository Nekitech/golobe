package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golobe/internal/database/model"
	"golobe/internal/services"
)

func HotelRoute(DB *sql.DB, router *gin.Engine) {
	var hotelMethods services.Hotel = &services.HotelScheme{
		DB:    nil,
		Hotel: model.Hotel{},
	}

	router.GET("/hotel", hotelMethods.GetHotels)
	router.GET("/hotel/:id", hotelMethods.GetHotelById)
	router.PATCH("/hotel/:id", hotelMethods.UpdateHotel)
	router.POST("/hotel", hotelMethods.CreateHotel)
	router.DELETE("/hotel/:id", hotelMethods.DeleteHotel)
}
