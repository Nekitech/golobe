package routes

import (
	"github.com/gin-gonic/gin"
	"golobe/controllers"
	"golobe/model"
	"gorm.io/gorm"
)

func HotelRoute(DB *gorm.DB, router *gin.Engine) {
	hotelMethods := controllers.HotelScheme{
		Model: gorm.Model{},
		DB:    DB,
		Hotel: model.Hotel{},
	}

	router.GET("/hotel", hotelMethods.GetHotels)
	router.GET("/hotel/:id", hotelMethods.GetHotelById)
	router.PATCH("/hotel/:id", hotelMethods.UpdateHotel)
	router.POST("/hotel", hotelMethods.CreateHotel)
	router.DELETE("/hotel/:id", hotelMethods.DeleteHotel)
}
