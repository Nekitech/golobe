package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golobe/internal/database/model"
	"golobe/internal/services"
)

func BookingRoute(DB *sql.DB, router *gin.Engine) {
	var bookingMethods services.Booking = &services.BookingScheme{
		DB:      DB,
		Booking: model.Booking{},
	}

	router.POST("/booking", bookingMethods.CreateBooking)
}
