package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golobe/internal/database/model"
	"golobe/internal/services"
)

func UserRoute(DB *sql.DB, router *gin.Engine) {
	var userMethods services.User = &services.UserScheme{
		DB:   DB,
		User: model.User{},
	}

	var bookingMethods services.Booking = &services.BookingScheme{
		DB:      DB,
		Booking: model.Booking{},
	}

	router.POST("/user", userMethods.CreateUser, bookingMethods.CreateUserHistoryBooking)
	router.PATCH("/user/:id", userMethods.UpdateUser)
}
