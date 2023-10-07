package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golobe/internal/controllers"
	"golobe/internal/model"
)

func BookingRoute(DB *sql.DB, router *gin.Engine) {
	bookingMethods := controllers.BookingScheme{
		DB:      DB,
		Booking: model.Booking{},
	}

	router.POST("/booking", bookingMethods.CreateBooking)
}
