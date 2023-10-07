package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	controllers2 "golobe/internal/controllers"
	model2 "golobe/internal/model"
)

func UserRoute(DB *sql.DB, router *gin.Engine) {
	userMethods := controllers2.UserScheme{
		DB:   DB,
		User: model2.User{},
	}

	bookingMethods := controllers2.BookingScheme{
		DB:      DB,
		Booking: model2.Booking{},
	}

	router.POST("/user", userMethods.CreateUser, bookingMethods.CreateUserHistoryBooking)
	router.PATCH("/user/:id", userMethods.UpdateUser)
}
